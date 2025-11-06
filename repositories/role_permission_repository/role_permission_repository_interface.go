package role_permission_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
)

type RolesPermissionsRepositoryInterface interface {
	GetDistinctPermissionByRoleId(ctx context.Context, tx *sqlx.Tx, roleId string) ([]models.GetPermissionName, *constants.ErrorResponse)
	UpsertRolePermission(ctx context.Context, tx *sqlx.Tx, data models.UpsertRolePermissionRequest) *constants.ErrorResponse
	DeleteRolePermission(ctx context.Context, tx *sqlx.Tx, data models.DeleteRolePermissionRequest) *constants.ErrorResponse
}

func NewRolesPermissionsRepository() RolesPermissionsRepositoryInterface {
	return &repository{}
}
