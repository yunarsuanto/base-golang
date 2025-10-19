package service

import (
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/services/activity_log_service"
	"github.com/yunarsuanto/base-go/services/auth_service"
)

type ServiceCtx struct {
	AuthService        auth_service.AuthServiceInterface
	ActivityLogService activity_log_service.ActivityLogServiceInterface
}

func InitServiceCtx(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) *ServiceCtx {
	return &ServiceCtx{
		AuthService:        auth_service.NewAuthService(repoCtx, infraCtx),
		ActivityLogService: activity_log_service.NewActivityLogService(repoCtx, infraCtx),
	}
}
