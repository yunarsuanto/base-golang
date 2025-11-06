package role_permission_repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/utils"
)

type repository struct{}

func (repository) GetDistinctPermissionByRoleId(ctx context.Context, tx *sqlx.Tx, roleId string) ([]models.GetPermissionName, *constants.ErrorResponse) {
	var result []models.GetPermissionName
	var query models.GetPermissionName

	params := []any{
		roleId,
	}

	err := tx.SelectContext(
		ctx,
		&result,
		fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), query.FilterQuery()),
		params...,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, constants.ErrIneligibleAccess
		}
		return result, utils.ErrDatabase(err, models.RolesPermissionsDataName)
	}

	return result, nil
}

func (repository) UpsertRolePermission(ctx context.Context, tx *sqlx.Tx, data models.UpsertRolePermissionRequest) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.RoleDataName)
	}

	return nil
}

func (repository) DeleteRolePermission(ctx context.Context, tx *sqlx.Tx, data models.DeleteRolePermissionRequest) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.RoleDataName)
	}

	return nil
}
