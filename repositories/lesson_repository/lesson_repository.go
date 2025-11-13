package lesson_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type repository struct{}

func (repository) ListLesson(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination, categoryLessonId string) ([]models.ListLesson, *constants.ErrorResponse) {
	var result []models.ListLesson
	var query models.ListLesson

	params := []any{}

	filterQuery, err := mapQueryFilterListLesson(pagination.Search, &params, categoryLessonId)
	if err != nil {
		return result, utils.ErrDatabase(err, models.LessonDatatitle)
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
		fmt.Println("kesini", err)
		return result, utils.ErrDatabase(err, models.LessonDatatitle)
	}

	var count int
	err = tx.QueryRowContext(
		ctx,
		countQuery,
		params...,
	).Scan(&count)
	if err != nil {
		return result, utils.ErrDatabase(err, models.LessonDatatitle)
	}

	pagination.GetPagination(count)
	return result, nil
}

func (repository) DetailLesson(ctx context.Context, tx *sqlx.Tx, id string) (models.DetailLesson, *constants.ErrorResponse) {
	var result []models.DetailLesson
	var query models.DetailLesson

	params := []any{id}

	getQuery := fmt.Sprintf("SELECT %s %s %s", query.ColumnQuery(), query.TableQuery(), query.FilterQuery())

	err := tx.SelectContext(
		ctx,
		&result,
		getQuery,
		params...,
	)
	if err != nil {
		return query, utils.ErrDatabase(err, models.LessonDatatitle)
	}

	if len(result) > 0 {
		query = result[0]
	}
	return query, nil
}

func (repository) CreateLesson(ctx context.Context, tx *sqlx.Tx, data models.CreateLesson) (string, *constants.ErrorResponse) {
	var id uuid.UUID

	query, args, err := sqlx.Named(data.InsertQuery(), data)
	if err != nil {
		return "", utils.ErrDatabase(err, models.LessonDatatitle)
	}

	query = tx.Rebind(query)

	utils.QueryLog(query, args...)
	err = tx.QueryRowxContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return "", utils.ErrDatabase(err, models.LessonDatatitle)
	}

	return id.String(), nil
}
func (repository) UpdateLesson(ctx context.Context, tx *sqlx.Tx, data models.UpdateLesson) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.LessonDatatitle)
	}

	return nil
}
func (repository) DeleteLesson(ctx context.Context, tx *sqlx.Tx, data models.DeleteLesson) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		data.InsertQuery(),
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.LessonDatatitle)
	}

	return nil
}
