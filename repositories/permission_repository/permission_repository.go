package permission_repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type repository struct{}

func (repository) ListPermission(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination) ([]models.ListPermission, *constants.ErrorResponse) {
	var result []models.ListPermission
	var query models.ListPermission

	params := []any{}

	filterQuery, err := mapQueryFilterListPermission(pagination.Search, &params)
	if err != nil {
		return result, utils.ErrDatabase(err, models.PermissionDataName)
	}

	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), filterQuery)
	countQuery := fmt.Sprintf("SELECT COUNT(1) %s %s", query.TableQuery(), filterQuery)
	if errs := utils.QueryOperation(&getQuery, [][2]string{{"u.name", constants.Ascending}}, pagination.Limit, pagination.Page); errs != nil {
		return result, errs
	}

	err = tx.SelectContext(
		ctx,
		&result,
		getQuery,
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.PermissionDataName)
	}

	var count int
	err = tx.QueryRowContext(
		ctx,
		countQuery,
		params...,
	).Scan(&count)
	if err != nil {
		return result, utils.ErrDatabase(err, models.PermissionDataName)
	}

	pagination.GetPagination(count)
	return result, nil
}

func (repository) CreatePermission(ctx context.Context, tx *sqlx.Tx, data models.CreatePermission) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.PermissionDataName)
	}

	return nil
}
func (repository) UpdatePermission(ctx context.Context, tx *sqlx.Tx, data models.UpdatePermission) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.PermissionDataName)
	}

	return nil
}
func (repository) DeletePermission(ctx context.Context, tx *sqlx.Tx, data models.DeletePermission) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.PermissionDataName)
	}

	return nil
}
