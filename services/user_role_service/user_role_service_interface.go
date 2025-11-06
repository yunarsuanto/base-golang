package user_role_service

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
)

type UserRoleServiceInterface interface {
	UpdateUserRole(ctx context.Context, req objects.UpsertUserRoleRequest) *constants.ErrorResponse
	DeleteUserRole(ctx context.Context, req objects.DeleteUserRoleRequest) *constants.ErrorResponse
}

func NewUserRoleService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) UserRoleServiceInterface {
	return &service{
		repoCtx,
		infraCtx,
	}
}
