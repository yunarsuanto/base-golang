package role_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

type RoleRepositoryInterface interface {
	ListRole(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination) ([]models.ListRole, *constants.ErrorResponse)
	DetailRole(ctx context.Context, tx *sqlx.Tx, id string) ([]models.DetailRole, *constants.ErrorResponse)
	CreateRole(ctx context.Context, tx *sqlx.Tx, data models.CreateRole) *constants.ErrorResponse
	UpdateRole(ctx context.Context, tx *sqlx.Tx, data models.UpdateRole) *constants.ErrorResponse
	DeleteRole(ctx context.Context, tx *sqlx.Tx, data models.DeleteRole) *constants.ErrorResponse
}

func NewRoleRepository() RoleRepositoryInterface {
	return &repository{}
}
