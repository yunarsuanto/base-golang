package role_permission_service

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
)

type RolePermissionServiceInterface interface {
	UpsertRolePermission(ctx context.Context, req objects.UpsertRolePermissionRequest) *constants.ErrorResponse
	DeleteRolePermission(ctx context.Context, req objects.DeleteRolePermissionRequest) *constants.ErrorResponse
}

func NewRolePermissionService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) RolePermissionServiceInterface {
	return &service{
		repoCtx,
		infraCtx,
	}
}
