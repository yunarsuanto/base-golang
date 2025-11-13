package lesson_item_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

type LessonItemRepositoryInterface interface {
	ListLessonItem(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination, lessonId string) ([]models.ListLessonItem, *constants.ErrorResponse)
	ListLessonItemByLessonIds(ctx context.Context, tx *sqlx.Tx, lessonIds []string) ([]models.ListLessonItem, *constants.ErrorResponse)
	DetailLessonItem(ctx context.Context, tx *sqlx.Tx, id string) (models.DetailLessonItem, *constants.ErrorResponse)
	CreateLessonItem(ctx context.Context, tx *sqlx.Tx, data models.CreateLessonItem) *constants.ErrorResponse
	UpdateLessonItem(ctx context.Context, tx *sqlx.Tx, data models.UpdateLessonItem) *constants.ErrorResponse
	DeleteLessonItem(ctx context.Context, tx *sqlx.Tx, data models.DeleteLessonItem) *constants.ErrorResponse

	BulkCreateLessonItem(ctx context.Context, tx *sqlx.Tx, data []models.BulkCreateLessonItem) *constants.ErrorResponse
}

func NewLessonItemRepository() LessonItemRepositoryInterface {
	return &repository{}
}
