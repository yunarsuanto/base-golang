package category_lesson_repository

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

func (repository) ListCategoryLesson(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination) ([]models.ListCategoryLesson, *constants.ErrorResponse) {
	var result []models.ListCategoryLesson
	var query models.ListCategoryLesson

	params := []any{}

	filterQuery, err := mapQueryFilterListCategoryLesson(pagination.Search, &params)
	if err != nil {
		return result, utils.ErrDatabase(err, models.CategoryLessonDataName)
	}

	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), filterQuery)
	countQuery := fmt.Sprintf("SELECT COUNT(1) %s %s", query.TableQuery(), filterQuery)
	if errs := utils.QueryOperation(&getQuery, [][2]string{{"u.title", constants.Ascending}}, pagination.Limit, pagination.Page); errs != nil {
		return result, errs
	}
	err = tx.SelectContext(
		ctx,
		&result,
		getQuery,
		params...,
	)
	if err != nil {
		return result, utils.ErrDatabase(err, models.CategoryLessonDataName)
	}

	var count int
	err = tx.QueryRowContext(
		ctx,
		countQuery,
		params...,
	).Scan(&count)
	if err != nil {
		return result, utils.ErrDatabase(err, models.CategoryLessonDataName)
	}

	pagination.GetPagination(count)
	return result, nil
}

func (repository) DetailCategoryLesson(ctx context.Context, tx *sqlx.Tx, id string) (models.DetailCategoryLesson, *constants.ErrorResponse) {
	var result []models.DetailCategoryLesson
	var query models.DetailCategoryLesson

	params := []any{id}

	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), query.FilterQuery())
	err := tx.SelectContext(
		ctx,
		&result,
		getQuery,
		params...,
	)
	if err != nil {
		return query, utils.ErrDatabase(err, models.CategoryLessonDataName)
	}

	if len(result) > 0 {
		query = result[0]
	}

	return query, nil
}

func (repository) CreateCategoryLesson(ctx context.Context, tx *sqlx.Tx, data models.CreateCategoryLesson) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.CategoryLessonDataName)
	}

	return nil
}
func (repository) UpdateCategoryLesson(ctx context.Context, tx *sqlx.Tx, data models.UpdateCategoryLesson) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.CategoryLessonDataName)
	}

	return nil
}
func (repository) DeleteCategoryLesson(ctx context.Context, tx *sqlx.Tx, data models.DeleteCategoryLesson) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.CategoryLessonDataName)
	}

	return nil
}

func (repository) CategoryLessonPublic(ctx context.Context, tx *sqlx.Tx) (models.ListCategoryLesson, *constants.ErrorResponse) {
	var result []models.ListCategoryLesson
	var query models.ListCategoryLesson

	params := []any{}

	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), "WHERE u.title ILIKE '%Mengenal Hewan%'")
	err := tx.SelectContext(
		ctx,
		&result,
		getQuery,
		params...,
	)
	if err != nil {
		return query, utils.ErrDatabase(err, models.CategoryLessonDataName)
	}

	if len(result) > 0 {
		query = result[0]
	}

	return query, nil
}
