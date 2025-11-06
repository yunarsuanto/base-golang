package user_repository

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

func (repository) GetByUsername(ctx context.Context, tx *sqlx.Tx, username string) (models.ListUser, *constants.ErrorResponse) {
	var result models.ListUser

	params := []any{
		username,
	}
	err := tx.GetContext(
		ctx,
		&result,
		fmt.Sprintf("SELECT %s %s WHERE (u.username = $1)", result.ColumnQuery(), result.TableQuery()),
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserDataName)
	}

	return result, nil
}

func (repository) GetById(ctx context.Context, tx *sqlx.Tx, id string) (models.ListUser, *constants.ErrorResponse) {
	var result models.ListUser

	params := []any{
		id,
	}
	err := tx.GetContext(
		ctx,
		&result,
		fmt.Sprintf("SELECT %s %s WHERE (u.id = $1)", result.ColumnQuery(), result.TableQuery()),
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserDataName)
	}

	return result, nil
}

func (repository) GetByTokenVerification(ctx context.Context, tx *sqlx.Tx, tokenVerification string) (models.ListUser, *constants.ErrorResponse) {
	var result models.ListUser

	params := []any{
		tokenVerification,
	}
	err := tx.GetContext(
		ctx,
		&result,
		fmt.Sprintf("SELECT %s %s WHERE (u.token_verification = $1)", result.ColumnQuery(), result.TableQuery()),
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserDataName)
	}

	return result, nil
}

func (repository) ListUser(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination) ([]models.ListUser, *constants.ErrorResponse) {
	var result []models.ListUser
	var query models.ListUser

	params := []any{}

	filterQuery, err := mapQueryFilterListUser(pagination.Search, &params)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserDataName)
	}

	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), filterQuery)
	countQuery := fmt.Sprintf("SELECT COUNT(1) %s %s", query.TableQuery(), filterQuery)
	if errs := utils.QueryOperation(&getQuery, [][2]string{{"u.username", constants.Ascending}}, pagination.Limit, pagination.Page); errs != nil {
		return result, errs
	}

	err = tx.SelectContext(
		ctx,
		&result,
		getQuery,
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserDataName)
	}

	var count int
	err = tx.QueryRowContext(
		ctx,
		countQuery,
		params...,
	).Scan(&count)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserDataName)
	}

	pagination.GetPagination(count)
	return result, nil
}

func (repository) DetailUser(ctx context.Context, tx *sqlx.Tx, id string) ([]models.DetailUser, *constants.ErrorResponse) {
	var result []models.DetailUser
	var query models.DetailUser

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

func (repository) CreateUser(ctx context.Context, tx *sqlx.Tx, data models.CreateUser) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UserDataName)
	}

	return nil
}
func (repository) UpdateUser(ctx context.Context, tx *sqlx.Tx, data models.UpdateUser) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UserDataName)
	}

	return nil
}
func (repository) DeleteUser(ctx context.Context, tx *sqlx.Tx, data models.DeleteUser) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UserDataName)
	}

	return nil
}
func (repository) UpdateTokenVerification(ctx context.Context, tx *sqlx.Tx, data models.UpdateUserTokenVerification) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UserDataName)
	}

	return nil
}
func (repository) UpdateTokenVerificationIsActiveUser(ctx context.Context, tx *sqlx.Tx, data models.UpdateUserIsActiveTokenVerification) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UserDataName)
	}

	return nil
}
func (repository) CreateUserProfile(ctx context.Context, tx *sqlx.Tx, data models.CreateUserProfile) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UserDataName)
	}

	return nil
}
