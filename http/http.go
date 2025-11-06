package http

import (
	"context"
	"database/sql"
	"embed"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/pressly/goose/v3"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/handler"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/infra/initiator/service"
	"github.com/yunarsuanto/base-go/infra/middleware"
	"github.com/yunarsuanto/base-go/infra/redis"
	"github.com/yunarsuanto/base-go/routers"
	"github.com/yunarsuanto/base-go/utils"

	config "github.com/yunarsuanto/base-go/config"
	db "github.com/yunarsuanto/base-go/infra/db"

	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

var EmbedMigration embed.FS

func migrate(dbAppConfig config.DBConfig) error {
	// Use migrations from the filesystem (constants.MigrationDir).
	// If you need to embed migrations, move the migrations directory into this package
	// and add a valid go:embed pattern that does not use "..".
	dbApp, err := sql.Open("postgres", dbAppConfig.Host)
	if err != nil {
		return err
	}
	defer dbApp.Close()
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(dbApp, constants.MigrationDir); err != nil {
		return err
	}

	return nil
}

func seed(ctx context.Context, repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) *constants.ErrorResponse {
	tx, err := infraCtx.Db.Begin(ctx)
	if err != nil {
		_ = tx.Rollback()
		return utils.ErrorInternalServer(err.Error())
	}

	// errs := repoCtx.ModuleRepo.Seed(ctx, tx)
	// if errs != nil {
	// 	return errs
	// }
	// errs = repoCtx.RolesModulesRepo.SeedSuperUser(ctx, tx)
	// if errs != nil {
	// 	return errs
	// }

	// errs = repoCtx.PermissionRepo.Seed(ctx, tx)
	// if errs != nil {
	// 	return errs
	// }
	// errs = repoCtx.RolesPermissionsRepo.SeedSuperUser(ctx, tx)
	// if errs != nil {
	// 	return errs
	// }

	// errs = repoCtx.NotificationTemplateRepo.Seed(ctx, tx)
	// if errs != nil {
	// 	return errs
	// }

	// errs = repoCtx.LocationRepo.Seed(ctx, tx)
	// if errs != nil {
	// 	return errs
	// }

	// errs = repoCtx.WorkShiftRepo.Seed(ctx, tx)
	// if errs != nil {
	// 	return errs
	// }

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return utils.ErrorInternalServer(err.Error())
	}

	return nil
}

func scheduler(serviceCtx *service.ServiceCtx) {
	tz, err := time.LoadLocation(constants.JakartaTimezone)
	if err != nil {
		errs := utils.ErrorInternalServer(err.Error())
		utils.PrintError(*errs)
		return
	}

	s := gocron.NewScheduler(tz)

	// _, err = s.Cron("* * * * *").Do(serviceCtx.CronService.CustomNotification)
	// if err != nil {
	// 	errs := utils.ErrorInternalServer(err.Error())
	// 	utils.PrintError(*errs)
	// }

	// _, err = s.Cron("0 0 * * *").Do(serviceCtx.CronService.SyncFormPeriod)
	// if err != nil {
	// 	errs := utils.ErrorInternalServer(err.Error())
	// 	utils.PrintError(*errs)
	// }

	// _, err = s.Cron("0 3 * * *").Do(serviceCtx.CronService.PurgeNotification)
	// if err != nil {
	// 	errs := utils.ErrorInternalServer(err.Error())
	// 	utils.PrintError(*errs)
	// }

	// _, err = s.Cron("0 6 * * *").Do(serviceCtx.CronService.EntryReminder)
	// if err != nil {
	// 	errs := utils.ErrorInternalServer(err.Error())
	// 	utils.PrintError(*errs)
	// }

	s.StartAsync()
}

func ServeHTTP() error {
	tz, err := time.LoadLocation(constants.JakartaTimezone)
	if err != nil {
		logrus.Fatalln(err)
	}

	// initial config
	cfg := config.InitConfig(tz)

	// init base context (used for startup tasks)
	baseCtx := context.Background()

	// database init
	dbApp, err := db.Open(&cfg.Db)
	if err != nil {
		logrus.Fatalln(err)
	}

	// db migration
	if err := migrate(cfg.Db); err != nil {
		logrus.Fatalln(err)
	}

	// Redis connection with timeout
	redisCtx, cancelRedis := context.WithTimeout(baseCtx, 5*time.Second)
	defer cancelRedis()

	redisServer := redis.NewRedisServer(&cfg.Redis)
	redisClient, err := redisServer.Connect(redisCtx)
	if err != nil {
		logrus.Fatalln(err)
	}

	// parse graceful shutdown timeout from CLI flag
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*time.Duration(cfg.Server.GraceFulTimeout), "duration to wait for active connections to finish")
	flag.Parse()

	// DI contexts
	repoCtx := repository.InitRepoCtx()
	infraCtx := infra.InitInfraCtx(dbApp, repoCtx, cfg, redisClient)
	serviceCtx := service.InitServiceCtx(repoCtx, infraCtx)

	// seeding with timeout
	seedCtx, cancelSeed := context.WithTimeout(baseCtx, 10*time.Second)
	defer cancelSeed()

	if errs := seed(seedCtx, repoCtx, infraCtx); errs != nil {
		logrus.Fatalln(errs.Err.Error())
	}

	// start scheduler in background
	go scheduler(serviceCtx)

	// setup CORS
	corsHandler := cors.New(cors.Options{
		AllowedHeaders: []string{
			constants.HeaderAuthorization,
			// constants.HeaderHospitalId,
			"Accept", "Origin", "Content-Type", "X-Requested-With", "Cache-Control",
		},
		AllowedMethods: []string{
			http.MethodHead, http.MethodGet, http.MethodPost,
			http.MethodPatch, http.MethodDelete, http.MethodOptions,
		},
		AllowedOrigins:     []string{"*"},
		ExposedHeaders:     []string{"Content-Disposition"},
		OptionsPassthrough: false,
		AllowCredentials:   true,
	})

	// middleware + router
	handlerCtx := handler.InitHandlerCtx(serviceCtx)
	mw := middleware.AccessTokenMiddleware(repoCtx, infraCtx)
	r := routers.InitRouter(handlerCtx, mw, &cfg)

	// configure HTTP server
	srv := &http.Server{
		Handler:      corsHandler.Handler(r),
		Addr:         cfg.Server.Addr,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
	}

	logrus.Println("API Listening on", cfg.Server.Addr)

	// start HTTP server in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Error("Server error:", err)
		}
	}()

	// handle graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM) // catch SIGINT & SIGTERM
	<-signalChan

	// shutdown context
	shutdownCtx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logrus.Error("Graceful shutdown failed:", err)
	}

	logrus.Println("Server shut down gracefully")
	return nil
}
