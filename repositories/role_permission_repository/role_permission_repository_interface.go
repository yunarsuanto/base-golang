package role_permission_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

type RolesPermissionsRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination, req objects.ListRolesPermissionsRequest) ([]models.GetRolesPermissions, *constants.ErrorResponse)
	GetByRoleIds(ctx context.Context, tx *sqlx.Tx, roleIds []string) ([]models.GetRolesPermissions, *constants.ErrorResponse)
	GetDistinctPermissionByRoleIds(ctx context.Context, tx *sqlx.Tx, roleIds []string) ([]models.GetPermissionCode, *constants.ErrorResponse)
	GetByRoleIdsPermissionCode(ctx context.Context, tx *sqlx.Tx, roleIds []string, permissionCode string) (models.GetRolesPermissions, *constants.ErrorResponse)
	Upsert(ctx context.Context, tx *sqlx.Tx, roleId string, permissionIds []string) *constants.ErrorResponse
	SeedSuperUser(ctx context.Context, tx *sqlx.Tx) *constants.ErrorResponse
}

func NewRolesPermissionsRepository() RolesPermissionsRepositoryInterface {
	return &rolesPermissionsRepository{}
}
