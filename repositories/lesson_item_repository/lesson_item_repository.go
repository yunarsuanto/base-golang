package lesson_item_repository

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

func (repository) ListLessonItem(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination) ([]models.ListLessonItem, *constants.ErrorResponse) {
	var result []models.ListLessonItem
	var query models.ListLessonItem

	params := []any{}

	filterQuery, err := mapQueryFilterListLessonItem(pagination.Search, &params)
	if err != nil {
		return result, utils.ErrDatabase(err, models.LessonItemDatacontent)
	}

	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), filterQuery)
	countQuery := fmt.Sprintf("SELECT COUNT(1) %s %s", query.TableQuery(), filterQuery)
	if errs := utils.QueryOperation(&getQuery, [][2]string{{"u.content", constants.Ascending}}, pagination.Limit, pagination.Page); errs != nil {
		return result, errs
	}

	err = tx.SelectContext(
		ctx,
		&result,
		getQuery,
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.LessonItemDatacontent)
	}

	var count int
	err = tx.QueryRowContext(
		ctx,
		countQuery,
		params...,
	).Scan(&count)
	if err != nil {
		return result, utils.ErrDatabase(err, models.LessonItemDatacontent)
	}

	pagination.GetPagination(count)
	return result, nil
}

func (repository) DetailLessonItem(ctx context.Context, tx *sqlx.Tx, id string) (models.DetailLessonItem, *constants.ErrorResponse) {
	var result []models.DetailLessonItem
	var query models.DetailLessonItem

	params := []any{id}

	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), query.FilterQuery())
	err := tx.SelectContext(
		ctx,
		&result,
		getQuery,
		params...,
	)
	if err != nil {
		return query, utils.ErrDatabase(err, models.LessonItemDatacontent)
	}

	if len(result) > 0 {
		query = result[0]
	}

	return query, nil
}

func (repository) CreateLessonItem(ctx context.Context, tx *sqlx.Tx, data models.CreateLessonItem) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.LessonItemDatacontent)
	}

	return nil
}
func (repository) UpdateLessonItem(ctx context.Context, tx *sqlx.Tx, data models.UpdateLessonItem) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.LessonItemDatacontent)
	}

	return nil
}
func (repository) DeleteLessonItem(ctx context.Context, tx *sqlx.Tx, data models.DeleteLessonItem) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.LessonItemDatacontent)
	}

	return nil
}
