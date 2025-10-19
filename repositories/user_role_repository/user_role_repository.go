package user_role_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type usersRolesRepository struct{}

func mapQueryFilterGetList(req objects.ListUsersRolesRequest, params *[]any) (string, error) {
	var filterArray []string
	var result string

	if len(req.RoleIds) != 0 {
		filterClause, filterArgs, err := sqlx.In("ur.role_id IN (?)", req.RoleIds)
		if err != nil {
			return result, err
		}
		filterArray = append(filterArray, filterClause)
		(*params) = append((*params), filterArgs...)
	}
	if len(req.UserIds) != 0 {
		filterClause, filterArgs, err := sqlx.In("ur.user_id IN (?)", req.UserIds)
		if err != nil {
			return result, err
		}
		filterArray = append(filterArray, filterClause)
		(*params) = append((*params), filterArgs...)
	}

	result = strings.Join(filterArray, " AND ")
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}

	return result, nil
}

func (usersRolesRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination, req objects.ListUsersRolesRequest) ([]models.GetUsersRoles, *constants.ErrorResponse) {
	var result []models.GetUsersRoles
	var query models.GetUsersRoles

	params := []any{}
	filterQuery, err := mapQueryFilterGetList(req, &params)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UsersRolesDataName)
	}
	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), filterQuery)
	countQuery := fmt.Sprintf("SELECT COUNT(1) %s %s", query.TableQuery(), filterQuery)
	if errs := utils.QueryOperation(&getQuery, [][2]string{{"r.name", constants.Ascending}}, pagination.Limit, pagination.Page); errs != nil {
		return result, errs
	}

	err = tx.SelectContext(
		ctx,
		&result,
		getQuery,
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UsersRolesDataName)
	}

	var count int
	err = tx.QueryRowContext(
		ctx,
		countQuery,
		params...,
	).Scan(&count)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UsersRolesDataName)
	}

	pagination.GetPagination(count)

	return result, nil
}

func (usersRolesRepository) Upsert(ctx context.Context, tx *sqlx.Tx, userId string, roleIds []string) *constants.ErrorResponse {
	deleteParams := []any{
		userId,
	}
	var additionalDeleteFilter string
	if len(roleIds) != 0 {
		deleteFilterClause, deleteFilterArgs, err := sqlx.In("role_id NOT IN (?)", roleIds)
		if err != nil {
			return utils.ErrDatabase(err, models.UsersRolesDataName)
		}
		additionalDeleteFilter = fmt.Sprintf("AND %s", deleteFilterClause)
		deleteParams = append(deleteParams, deleteFilterArgs...)
	}

	_, err := tx.ExecContext(
		ctx,
		fmt.Sprintf(deleteQuery, additionalDeleteFilter),
		deleteParams...,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UsersRolesDataName)
	}

	var upsertData []models.CreateUsersRoles
	for _, v := range roleIds {
		upsertData = append(upsertData, models.CreateUsersRoles{
			UserId: userId,
			RoleId: v,
		})
	}

	if len(upsertData) != 0 {
		_, err := tx.NamedExecContext(
			ctx,
			upsertQuery,
			upsertData,
		)
		if err != nil {
			return utils.ErrDatabase(err, models.UsersRolesDataName)
		}
	}

	return nil
}

func (usersRolesRepository) GetByPermissions(ctx context.Context, tx *sqlx.Tx, permissions []string) ([]models.GetUsersPermissions, *constants.ErrorResponse) {
	var result []models.GetUsersPermissions
	var query models.GetUsersPermissions

	filterQuery, filterArgs, err := sqlx.In("p.code IN (?)", permissions)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UsersRolesDataName)
	}

	err = tx.SelectContext(
		ctx,
		&result,
		fmt.Sprintf("SELECT %s %s WHERE %s", query.ColumnQuery(), query.TableQuery(), filterQuery),
		filterArgs...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UsersRolesDataName)
	}

	return result, nil
}
