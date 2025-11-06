package user_role_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

type UsersRolesRepositoryInterface interface {
	GetByUserId(ctx context.Context, tx *sqlx.Tx, req objects.ListUserRoleRequest) (models.GetUserRole, *constants.ErrorResponse)
	UpsertUserRole(ctx context.Context, tx *sqlx.Tx, data models.UpsertUserRoleRequest) *constants.ErrorResponse
	DeleteUserRole(ctx context.Context, tx *sqlx.Tx, data models.DeleteUserRoleRequest) *constants.ErrorResponse
}

func NewUsersRolesRepository() UsersRolesRepositoryInterface {
	return &repository{}
}
