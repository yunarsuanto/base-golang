package infra

import (
	"github.com/yunarsuanto/base-go/config"
	"github.com/yunarsuanto/base-go/infra/db"
	"github.com/yunarsuanto/base-go/infra/file"
	"github.com/yunarsuanto/base-go/infra/firebase"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	jwt "github.com/yunarsuanto/base-go/infra/jwt"
	"github.com/yunarsuanto/base-go/infra/mail"

	redis "github.com/go-redis/redis/v8"
)

type InfraCtx struct {
	Db          *db.DB
	Config      *config.Config
	RedisClient *redis.Client
	Jwt         jwt.JWTInterface
	Storage     file.FileCtx
	Mailer      mail.MailInterface
	Firebase    firebase.FirebaseInterface
}

func InitInfraCtx(dbApp *db.DB, repo *repository.RepoCtx, cfg config.Config, rdb *redis.Client) *InfraCtx {
	jwtSvc := jwt.NewJWT(&cfg.JwtConfig, rdb, false)

	localStorage := file.NewLocalStorage(&cfg)
	urlStorage := file.NewUrlStorage(&cfg)

	fileSvc := file.FileCtx{
		Config: cfg.Server,
		Local:  localStorage,
		Url:    urlStorage,
	}

	return &InfraCtx{
		Db:          dbApp,
		Config:      &cfg,
		RedisClient: rdb,
		Jwt:         jwtSvc,
		Storage:     fileSvc,
		Mailer:      mail.NewMailService(dbApp, repo, &cfg),
		Firebase:    firebase.NewFirebaseService(dbApp, repo, &cfg),
	}
}
