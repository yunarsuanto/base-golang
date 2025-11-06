package permission_service

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
)

type PermissionServiceInterface interface {
	ListPermission(ctx context.Context, pagination *objects.Pagination) ([]objects.ListPermissionResponse, *constants.ErrorResponse)
	CreatePermission(ctx context.Context, req objects.CreatePermissionRequest) *constants.ErrorResponse
	UpdatePermission(ctx context.Context, req objects.UpdatePermissionRequest) *constants.ErrorResponse
	DeletePermission(ctx context.Context, req objects.DeletePermissionRequest) *constants.ErrorResponse
}

func NewPermissionService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) PermissionServiceInterface {
	return &service{
		repoCtx,
		infraCtx,
	}
}
