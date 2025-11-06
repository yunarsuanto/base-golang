package permission_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

type PermissionRepositoryInterface interface {
	ListPermission(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination) ([]models.ListPermission, *constants.ErrorResponse)
	CreatePermission(ctx context.Context, tx *sqlx.Tx, data models.CreatePermission) *constants.ErrorResponse
	UpdatePermission(ctx context.Context, tx *sqlx.Tx, data models.UpdatePermission) *constants.ErrorResponse
	DeletePermission(ctx context.Context, tx *sqlx.Tx, data models.DeletePermission) *constants.ErrorResponse
}

func NewPermissionRepository() PermissionRepositoryInterface {
	return &repository{}
}
