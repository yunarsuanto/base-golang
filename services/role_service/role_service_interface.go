package role_service

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
)

type RoleServiceInterface interface {
	ListRole(ctx context.Context, pagination *objects.Pagination) ([]objects.ListRoleResponse, *constants.ErrorResponse)
	DetailRole(ctx context.Context, req objects.DetailRoleRequest) ([]objects.DetailRoleResponse, *constants.ErrorResponse)
	CreateRole(ctx context.Context, req objects.CreateRoleRequest) *constants.ErrorResponse
	UpdateRole(ctx context.Context, req objects.UpdateRoleRequest) *constants.ErrorResponse
	DeleteRole(ctx context.Context, req objects.DeleteRoleRequest) *constants.ErrorResponse
}

func NewRoleService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) RoleServiceInterface {
	return &service{
		repoCtx,
		infraCtx,
	}
}
