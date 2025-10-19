package user_token_repository

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

type userTokenRepository struct{}

func mapQueryFilterGetList(req objects.ListUserTokenRequest, params *[]any) (string, error) {
	filterArray := []string{
		"ut.fcm_is_valid IS true",
		"ut.expiry_time > CURRENT_TIMESTAMP",
	}
	var result string

	if len(req.UserIds) != 0 {
		filterClause, filterArgs, err := sqlx.In("ut.user_id IN (?)", req.UserIds)
		if err != nil {
			return result, err
		}
		filterArray = append(filterArray, filterClause)
		(*params) = append((*params), filterArgs...)
	}
	if req.Platform != "" {
		filterArray = append(filterArray, "ut.platform = ?")
		(*params) = append((*params), req.Platform)
	}
	if req.GetFcmToken != nil {
		filterArray = append(filterArray, "(ut.fcm_token != '') = ?")
		(*params) = append((*params), *req.GetFcmToken)
	}

	result = strings.Join(filterArray, " AND ")
	if result != "" {
		result = fmt.Sprintf("WHERE %s", result)
	}

	return result, nil
}

func (userTokenRepository) GetList(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination, req objects.ListUserTokenRequest) ([]models.GetUserToken, *constants.ErrorResponse) {
	var result []models.GetUserToken
	var query models.GetUserToken

	params := []any{}
	filterQuery, err := mapQueryFilterGetList(req, &params)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserTokenDataName)
	}
	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), filterQuery)
	countQuery := fmt.Sprintf("SELECT COUNT(1) %s %s", query.TableQuery(), filterQuery)
	if errs := utils.QueryOperation(&getQuery, nil, pagination.Limit, pagination.Page); errs != nil {
		return result, errs
	}

	err = tx.SelectContext(
		ctx,
		&result,
		getQuery,
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserTokenDataName)
	}

	var count int
	err = tx.QueryRowContext(
		ctx,
		countQuery,
		params...,
	).Scan(&count)
	if err != nil {
		return result, utils.ErrDatabase(err, models.UserTokenDataName)
	}

	pagination.GetPagination(count)

	return result, nil
}

func (userTokenRepository) Upsert(ctx context.Context, tx *sqlx.Tx, data models.UpsertUserToken) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		upsertQuery,
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UserTokenDataName)
	}

	return nil
}

func (userTokenRepository) Delete(ctx context.Context, tx *sqlx.Tx, userId, platform string) *constants.ErrorResponse {
	_, err := tx.ExecContext(
		ctx,
		deleteQuery,
		userId,
		platform,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.UserTokenDataName)
	}

	return nil
}
