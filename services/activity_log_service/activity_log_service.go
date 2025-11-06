package activity_log_service

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type activityLogService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (a activityLogService) Create(ctx context.Context, data objects.CreateActivityLog) *constants.ErrorResponse {
	r := data.Request
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	var userId sql.NullString
	if claims, errs := utils.GetJwtClaimsFromContext(ctx); errs == nil {
		userId = utils.NewNullString(claims.Id)
	}

	var body sql.NullString
	if data.Body != nil {
		jsonBody, err := json.Marshal(data.Body)
		if err != nil {
			_ = tx.Rollback()
			return utils.ErrorInternalServer(err.Error())
		}

		body = utils.NewNullString(string(jsonBody))
	}

	var errorMessage sql.NullString
	if utils.FirstDigit(data.ResponseMeta.Status) != 2 {
		if data.ResponseMeta.Message == "" {
			errorMessage = utils.NewNullString("empty")
		} else {
			errorMessage = utils.NewNullString(data.ResponseMeta.Message)
		}
	} else {
		errorMessage = utils.NewNullString(data.ResponseMeta.Message)
	}
	createData := models.CreateActivityLog{
		UserId:       userId,
		Host:         r.Host,
		Path:         r.URL.Path,
		Body:         body,
		StatusCode:   data.ResponseMeta.Status,
		ErrorMessage: errorMessage,
		IpAddress:    utils.GetIP(r),
		UserAgent:    r.UserAgent(),
		MemoryUsage:  utils.GetMemoryUsage(),
	}
	errs := a.ActivityLogRepo.Create(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return utils.ErrorInternalServer(err.Error())
	}

	return nil
}
