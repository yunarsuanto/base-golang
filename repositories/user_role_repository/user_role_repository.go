package user_role_repository

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

func (repository) GetByUserId(ctx context.Context, tx *sqlx.Tx, req objects.ListUserRoleRequest) (models.GetUserRole, *constants.ErrorResponse) {
	var resModel []models.GetUserRole
	var result models.GetUserRole
	var query models.GetUserRole

	params := []any{
		req.UserId,
	}
	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), query.FilterQuery())

	err := tx.SelectContext(
		ctx,
		&resModel,
		getQuery,
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UsersRolesDataName)
	}

	if len(resModel) > 0 {
		result = resModel[0]
	}

	return result, nil
}
func (repository) UpsertUserRole(ctx context.Context, tx *sqlx.Tx, data models.UpsertUserRoleRequest) *constants.ErrorResponse {
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
func (repository) DeleteUserRole(ctx context.Context, tx *sqlx.Tx, data models.DeleteUserRoleRequest) *constants.ErrorResponse {
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
