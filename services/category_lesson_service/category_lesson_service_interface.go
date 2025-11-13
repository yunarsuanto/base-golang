package category_lesson_service

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
)

type CategoryLessonServiceInterface interface {
	ListCategoryLesson(ctx context.Context, pagination *objects.Pagination) ([]objects.ListCategoryLessonResponse, *constants.ErrorResponse)
	DetailCategoryLesson(ctx context.Context, req objects.DetailCategoryLessonRequest) (objects.DetailCategoryLessonResponse, *constants.ErrorResponse)
	CreateCategoryLesson(ctx context.Context, req objects.CreateCategoryLessonRequest) *constants.ErrorResponse
	UpdateCategoryLesson(ctx context.Context, req objects.UpdateCategoryLessonRequest) *constants.ErrorResponse
	DeleteCategoryLesson(ctx context.Context, req objects.DeleteCategoryLessonRequest) *constants.ErrorResponse

	CategoryLessonPublic(ctx context.Context) (objects.ListCategoryLessonPublicResponse, *constants.ErrorResponse)
}

func NewCategoryLessonService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) CategoryLessonServiceInterface {
	return &service{
		repoCtx,
		infraCtx,
	}
}
