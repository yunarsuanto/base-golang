package lesson_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

type LessonRepositoryInterface interface {
	ListLesson(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination, categoryLessonId string) ([]models.ListLesson, *constants.ErrorResponse)
	DetailLesson(ctx context.Context, tx *sqlx.Tx, id string) (models.DetailLesson, *constants.ErrorResponse)
	CreateLesson(ctx context.Context, tx *sqlx.Tx, data models.CreateLesson) (string, *constants.ErrorResponse)
	UpdateLesson(ctx context.Context, tx *sqlx.Tx, data models.UpdateLesson) *constants.ErrorResponse
	DeleteLesson(ctx context.Context, tx *sqlx.Tx, data models.DeleteLesson) *constants.ErrorResponse
}

func NewLessonRepository() LessonRepositoryInterface {
	return &repository{}
}
