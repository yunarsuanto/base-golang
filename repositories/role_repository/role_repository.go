package role_repository

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

func (repository) ListRole(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination) ([]models.ListRole, *constants.ErrorResponse) {
	var result []models.ListRole
	var query models.ListRole

	params := []any{}

	filterQuery, err := mapQueryFilterListRole(pagination.Search, &params)
	if err != nil {
		return result, utils.ErrDatabase(err, models.RoleDataName)
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
		return result, utils.ErrDatabase(err, models.RoleDataName)
	}

	var count int
	err = tx.QueryRowContext(
		ctx,
		countQuery,
		params...,
	).Scan(&count)
	if err != nil {
		return result, utils.ErrDatabase(err, models.RoleDataName)
	}

	pagination.GetPagination(count)
	return result, nil
}

func (repository) DetailRole(ctx context.Context, tx *sqlx.Tx, id string) ([]models.DetailRole, *constants.ErrorResponse) {
	var result []models.DetailRole
	var query models.DetailRole

	params := []any{id}

	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), query.FilterQuery())

	err := tx.SelectContext(
		ctx,
		&result,
		getQuery,
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserDataName)
	}

	return result, nil
}

func (repository) CreateRole(ctx context.Context, tx *sqlx.Tx, data models.CreateRole) *constants.ErrorResponse {
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
func (repository) UpdateRole(ctx context.Context, tx *sqlx.Tx, data models.UpdateRole) *constants.ErrorResponse {
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
func (repository) DeleteRole(ctx context.Context, tx *sqlx.Tx, data models.DeleteRole) *constants.ErrorResponse {
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
