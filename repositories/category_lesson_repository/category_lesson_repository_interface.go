package category_lesson_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

type CategoryLessonRepositoryInterface interface {
	ListCategoryLesson(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination, hasChild bool) ([]models.ListCategoryLesson, *constants.ErrorResponse)
	DetailCategoryLesson(ctx context.Context, tx *sqlx.Tx, id string) ([]models.DetailCategoryLesson, *constants.ErrorResponse)
	CreateCategoryLesson(ctx context.Context, tx *sqlx.Tx, data models.CreateCategoryLesson) *constants.ErrorResponse
	UpdateCategoryLesson(ctx context.Context, tx *sqlx.Tx, data models.UpdateCategoryLesson) *constants.ErrorResponse
	DeleteCategoryLesson(ctx context.Context, tx *sqlx.Tx, data models.DeleteCategoryLesson) *constants.ErrorResponse
}

func NewCategoryLessonRepository() CategoryLessonRepositoryInterface {
	return &repository{}
}
