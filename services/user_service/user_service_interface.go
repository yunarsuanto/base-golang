package user_service

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
)

type UserServiceInterface interface {
	ListUser(ctx context.Context, pagination *objects.Pagination) ([]objects.ListUserResponse, *constants.ErrorResponse)
	DetailUser(ctx context.Context, req objects.DetailUserRequest) ([]objects.DetailUserResponse, *constants.ErrorResponse)
	CreateUser(ctx context.Context, req objects.CreateUserRequest) *constants.ErrorResponse
	UpdateUser(ctx context.Context, req objects.UpdateUserRequest) *constants.ErrorResponse
	DeleteUser(ctx context.Context, req objects.DeleteUserRequest) *constants.ErrorResponse
}

func NewUserService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) UserServiceInterface {
	return &service{
		repoCtx,
		infraCtx,
	}
}
