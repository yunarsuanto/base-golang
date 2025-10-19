package user_role_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

type UsersRolesRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination, req objects.ListUsersRolesRequest) ([]models.GetUsersRoles, *constants.ErrorResponse)
	Upsert(ctx context.Context, tx *sqlx.Tx, uesrId string, roleIds []string) *constants.ErrorResponse
	GetByPermissions(ctx context.Context, tx *sqlx.Tx, permissions []string) ([]models.GetUsersPermissions, *constants.ErrorResponse)
}

func NewUsersRolesRepository() UsersRolesRepositoryInterface {
	return &usersRolesRepository{}
}
