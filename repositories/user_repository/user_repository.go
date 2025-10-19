package user_repository

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

type userRepository struct{}

func mapQueryFilterGetList(search string, req objects.ListUserRequest, params *[]any) string {
	var filterArray []string

	if search != "" {
		searchParams := fmt.Sprintf("%%%s%%", search)
		filterArray = append(filterArray, "u.email LIKE ? OR u.username LIKE ? OR u.name LIKE ?")
		(*params) = append((*params), searchParams, searchParams, searchParams)
	}
	if req.IsReporter != nil {
		filterArray = append(filterArray, "ur.is_reporter = ?")
		(*params) = append((*params), *req.IsReporter)
	}
	if req.IsVerificator != nil {
		filterArray = append(filterArray, "ur.is_verificator = ?")
		(*params) = append((*params), *req.IsVerificator)
	}
	if req.IsActive != nil {
		filterArray = append(filterArray, "u.is_active = ?")
		(*params) = append((*params), *req.IsActive)
	}

	result := strings.Join(filterArray, " AND ")
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}

	return result
}

func (userRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination, req objects.ListUserRequest) ([]models.GetUser, *constants.ErrorResponse) {
	var result []models.GetUser
	var query models.GetUser

	params := []any{}
	filterQuery := mapQueryFilterGetList(pagination.Search, req, &params)
	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), filterQuery)
	countQuery := fmt.Sprintf("SELECT COUNT(1) %s %s", query.TableQuery(), filterQuery)
	if errs := utils.QueryOperation(&getQuery, [][2]string{{"u.name", constants.Ascending}}, pagination.Limit, pagination.Page); errs != nil {
		return result, errs
	}

	err := tx.SelectContext(
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

func (userRepository) GetById(ctx context.Context, tx *sqlx.Tx, userId string) (models.GetUser, *constants.ErrorResponse) {
	var result models.GetUser

	err := tx.GetContext(
		ctx,
		&result,
		fmt.Sprintf("SELECT %s %s WHERE u.id = ?", result.ColumnQuery(), result.TableQuery()),
		userId,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserDataName)
	}

	return result, nil
}

func (userRepository) GetByEmployeeId(ctx context.Context, tx *sqlx.Tx, employeeId string) (models.GetUser, *constants.ErrorResponse) {
	var result models.GetUser

	err := tx.GetContext(
		ctx,
		&result,
		fmt.Sprintf("SELECT %s %s WHERE u.employee_id = ?", result.ColumnQuery(), result.TableQuery()),
		employeeId,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserDataName)
	}

	return result, nil
}

func (userRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateUser) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		createQuery,
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UserDataName)
	}

	return nil
}

func (userRepository) Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateUser) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		updateQuery,
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UserDataName)
	}

	return nil
}

func (userRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteQuery,
		id,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UserDataName)
	}

	return nil
}

func (userRepository) GetByEmail(ctx context.Context, tx *sqlx.Tx, email, username, excludingId string) (models.GetUser, *constants.ErrorResponse) {
	var result models.GetUser

	params := []any{
		email,
	}
	if username != "" {
		params = append(params, username)
	} else {
		params = append(params, email)
	}

	var additionalFilter string
	if excludingId != "" {
		additionalFilter = `AND u.id != ? AND u.employee_id != ?`
		params = append(params, excludingId, excludingId)
	}

	err := tx.GetContext(
		ctx,
		&result,
		fmt.Sprintf("SELECT %s %s WHERE (u.email LIKE ? OR u.username LIKE ?) %s", result.ColumnQuery(), result.TableQuery(), additionalFilter),
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserDataName)
	}

	return result, nil
}

func (userRepository) GetByIdNumber(ctx context.Context, tx *sqlx.Tx, idNumber, excludingId string) (models.GetUser, *constants.ErrorResponse) {
	var result models.GetUser

	params := []any{
		idNumber,
	}

	var additionalFilter string
	if excludingId != "" {
		additionalFilter = `AND u.id != ? AND u.employee_id != ?`
		params = append(params, excludingId, excludingId)
	}

	err := tx.GetContext(
		ctx,
		&result,
		fmt.Sprintf("SELECT %s %s WHERE u.employee_number = ? %s", result.ColumnQuery(), result.TableQuery(), additionalFilter),
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserDataName)
	}

	return result, nil
}

func (userRepository) GetByPhoneNumber(ctx context.Context, tx *sqlx.Tx, phoneNumber, excludingId string) (models.GetUser, *constants.ErrorResponse) {
	var result models.GetUser

	params := []any{
		phoneNumber,
	}

	var additionalFilter string
	if excludingId != "" {
		additionalFilter = `AND u.id != ? AND u.employee_id != ?`
		params = append(params, excludingId, excludingId)
	}

	err := tx.GetContext(
		ctx,
		&result,
		fmt.Sprintf("SELECT %s %s WHERE u.phone_number = ? %s", result.ColumnQuery(), result.TableQuery(), additionalFilter),
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserDataName)
	}

	return result, nil
}

func (userRepository) ChangePassword(ctx context.Context, tx *sqlx.Tx, data models.ChangeUserPassword) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		changePasswordQuery,
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UserDataName)
	}

	return nil
}

func (userRepository) UpdateProfile(ctx context.Context, tx *sqlx.Tx, data models.UpdateUser) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		updateProfileQuery,
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UserDataName)
	}

	return nil
}

func (userRepository) UpdateActivation(ctx context.Context, tx *sqlx.Tx, id string, isActive bool) *constants.ErrorResponse {
	// _, err := tx.ExecContext(
	// 	ctx,
	// 	updateActivationQuery,
	// 	isActive,
	// 	id,
	// )
	// if err != nil {
	// 	return utils.ErrDatabase(err, models.ClusterDataName)
	// }

	return nil
}
