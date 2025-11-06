package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/yunarsuanto/base-go/config"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/db"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"

	"google.golang.org/api/option"
)

type FirebaseInterface interface {
	BulkSend(req objects.SendNotification) *constants.ErrorResponse
}

type firebaseService struct {
	*db.DB
	*repository.RepoCtx
	*config.Config
}

func NewFirebaseService(dbApp *db.DB, repo *repository.RepoCtx, config *config.Config) FirebaseInterface {
	return &firebaseService{
		dbApp,
		repo,
		config,
	}
}

func (a firebaseService) initFirebase() (*firebase.App, *constants.ErrorResponse) {
	opt := option.WithCredentialsFile(a.Firebase.ServiceKeyPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, utils.ErrorInternalServer(err.Error())
	}

	return app, nil
}

func (a firebaseService) BulkSend(req objects.SendNotification) *constants.ErrorResponse {
	// ctx := context.Background()
	// defer ctx.Done()

	// app, errs := a.initFirebase()
	// if errs != nil {
	// 	return errs
	// }

	// client, err := app.Messaging(ctx)
	// if err != nil {
	// 	return utils.ErrorInternalServer(err.Error())
	// }

	// tx, err := a.Begin(ctx)
	// if err != nil {
	// 	return utils.ErrorInternalServer(err.Error())
	// }

	// getFcmToken := true
	// tokenData, errs := a.UserTokenRepo.GetList(ctx, tx, objects.NewPagination().AllData(), objects.ListUserTokenRequest{
	// 	UserIds:     req.UserIds,
	// 	GetFcmToken: &getFcmToken,
	// })
	// if errs != nil {
	// 	_ = tx.Rollback()
	// 	return errs
	// }

	// var data []*messaging.Message
	// for _, v := range tokenData {
	// 	data = append(data, &messaging.Message{
	// 		Notification: &messaging.Notification{
	// 			Title: req.Title,
	// 			Body:  req.Body,
	// 		},
	// 		Token: v.FcmToken,
	// 	})
	// }

	// for i := 0; i < len(data); i += constants.MaximumSendBatch {
	// 	lastIndex := i + constants.MaximumSendBatch
	// 	if lastIndex > len(data) {
	// 		lastIndex = len(data)
	// 	}

	// 	go func() {
	// 		res, err := client.SendEach(ctx, data[i:lastIndex])
	// 		if err != nil {
	// 			logrus.Errorln("Error Send Notification:", err.Error())
	// 		}
	// 		if res.FailureCount != 0 {
	// 			for _, v := range res.Responses {
	// 				if v.Error != nil {
	// 					logrus.Errorln("Send Notification Failed:", v.Error.Error())
	// 				}
	// 			}
	// 		}
	// 	}()
	// }

	// err = tx.Commit()
	// if err != nil {
	// 	_ = tx.Rollback()
	// 	return utils.ErrorInternalServer(err.Error())
	// }

	return nil
}
