package role_permission_repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type rolesPermissionsRepository struct{}

func mapQueryFilterGetList(req objects.ListRolesPermissionsRequest, params *[]any) string {
	var filterArray []string

	if req.RoleId != "" {
		filterArray = append(filterArray, "rp.role_id = ?")
		(*params) = append((*params), req.RoleId)
	}

	result := strings.Join(filterArray, " AND ")
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}

	return result
}

func (rolesPermissionsRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination, req objects.ListRolesPermissionsRequest) ([]models.GetRolesPermissions, *constants.ErrorResponse) {
	var result []models.GetRolesPermissions
	var query models.GetRolesPermissions

	params := []any{}
	filterQuery := mapQueryFilterGetList(req, &params)
	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), filterQuery)
	countQuery := fmt.Sprintf("SELECT COUNT(1) %s %s", query.TableQuery(), filterQuery)
	if errs := utils.QueryOperation(&getQuery, [][2]string{{"r.name", constants.Ascending}, {"p.name", constants.Ascending}}, pagination.Limit, pagination.Page); errs != nil {
		return result, errs
	}

	err := tx.SelectContext(
		ctx,
		&result,
		getQuery,
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.RolesPermissionsDataName)
	}

	var count int
	err = tx.QueryRowContext(
		ctx,
		countQuery,
		params...,
	).Scan(&count)
	if err != nil {
		return result, utils.ErrDatabase(err, models.RolesPermissionsDataName)
	}

	pagination.GetPagination(count)

	return result, nil
}

func (rolesPermissionsRepository) GetByRoleIds(ctx context.Context, tx *sqlx.Tx, roleIds []string) ([]models.GetRolesPermissions, *constants.ErrorResponse) {
	var result []models.GetRolesPermissions
	var query models.GetRolesPermissions

	if len(roleIds) == 0 {
		return result, nil
	}
	filterClause, filterArgs, err := sqlx.In("rp.role_id IN (?)", roleIds)
	if err != nil {
		return result, utils.ErrDatabase(err, models.RolesPermissionsDataName)
	}

	err = tx.SelectContext(
		ctx,
		&result,
		fmt.Sprintf("SELECT %s %s WHERE %s ORDER BY p.code ASC", query.ColumnQuery(), query.TableQuery(), filterClause),
		filterArgs...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.RolesPermissionsDataName)
	}

	return result, nil
}

func (rolesPermissionsRepository) GetByRoleIdsPermissionCode(ctx context.Context, tx *sqlx.Tx, roleIds []string, permissionCode string) (models.GetRolesPermissions, *constants.ErrorResponse) {
	var result models.GetRolesPermissions

	params := []any{permissionCode}
	filterClause, filterArgs, err := sqlx.In("r.id IN (?)", roleIds)
	if err != nil {
		return result, utils.ErrDatabase(err, models.RolesPermissionsDataName)
	}
	params = append(params, filterArgs...)

	err = tx.GetContext(
		ctx,
		&result,
		fmt.Sprintf("SELECT %s %s WHERE p.code = ? AND %s", result.ColumnQuery(), result.TableQuery(), filterClause),
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

func (rolesPermissionsRepository) GetDistinctPermissionByRoleIds(ctx context.Context, tx *sqlx.Tx, roleIds []string) ([]models.GetPermissionCode, *constants.ErrorResponse) {
	var result []models.GetPermissionCode
	var query models.GetPermissionCode

	if len(roleIds) == 0 {
		return result, nil
	}

	filterClause, filterArgs, err := sqlx.In("rp.role_id IN (?)", roleIds)
	if err != nil {
		return result, utils.ErrDatabase(err, models.RolesPermissionsDataName)
	}

	err = tx.SelectContext(
		ctx,
		&result,
		fmt.Sprintf("SELECT %s %s WHERE %s ORDER BY permission_code", query.ColumnQuery(), query.TableQuery(), filterClause),
		filterArgs...,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, constants.ErrIneligibleAccess
		}
		return result, utils.ErrDatabase(err, models.RolesPermissionsDataName)
	}

	return result, nil
}

func (rolesPermissionsRepository) Upsert(ctx context.Context, tx *sqlx.Tx, roleId string, permissionIds []string) *constants.ErrorResponse {
	deleteParams := []any{
		roleId,
	}
	var additionalDeleteFilter string
	if len(permissionIds) != 0 {
		filterClause, filterArgs, err := sqlx.In("permission_id NOT IN (?)", permissionIds)
		if err != nil {
			return utils.ErrDatabase(err, models.RolesPermissionsDataName)
		}
		additionalDeleteFilter = fmt.Sprintf("AND %s", filterClause)
		deleteParams = append(deleteParams, filterArgs...)
	}

	_, err := tx.ExecContext(
		ctx,
		fmt.Sprintf(deleteQuery, additionalDeleteFilter),
		deleteParams...,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.RolesPermissionsDataName)
	}

	var upsertData []models.CreateRolesPermissions
	for _, v := range permissionIds {
		upsertData = append(upsertData, models.CreateRolesPermissions{
			RoleId:       roleId,
			PermissionId: v,
		})
	}

	if len(upsertData) != 0 {
		_, err := tx.NamedExecContext(
			ctx,
			upsertQuery,
			upsertData,
		)
		if err != nil {
			return utils.ErrDatabase(err, models.RolesPermissionsDataName)
		}
	}

	return nil
}

func (rolesPermissionsRepository) SeedSuperUser(ctx context.Context, tx *sqlx.Tx) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		seedSuperUserQuery,
		constants.SuperUserRoleName,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.RolesPermissionsDataName)
	}

	return nil
}
