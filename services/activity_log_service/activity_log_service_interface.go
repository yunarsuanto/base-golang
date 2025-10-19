package activity_log_service

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
)

type ActivityLogServiceInterface interface {
	Create(ctx context.Context, data objects.CreateActivityLog) *constants.ErrorResponse
}

func NewActivityLogService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) ActivityLogServiceInterface {
	return &activityLogService{
		repoCtx,
		infraCtx,
	}
}
