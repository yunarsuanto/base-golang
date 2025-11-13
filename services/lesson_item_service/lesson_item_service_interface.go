package lesson_item_service

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
)

type LessonItemServiceInterface interface {
	ListLessonItem(ctx context.Context, pagination *objects.Pagination, lessonId string) ([]objects.ListLessonItemResponse, *constants.ErrorResponse)
	ListLessonItemByLessonIds(ctx context.Context, lessonId []string) ([]objects.ListLessonItemResponse, *constants.ErrorResponse)

	DetailLessonItem(ctx context.Context, req objects.DetailLessonItemRequest) (objects.DetailLessonItemResponse, *constants.ErrorResponse)
	CreateLessonItem(ctx context.Context, req objects.CreateLessonItemRequest) *constants.ErrorResponse
	UpdateLessonItem(ctx context.Context, req objects.UpdateLessonItemRequest) *constants.ErrorResponse
	DeleteLessonItem(ctx context.Context, req objects.DeleteLessonItemRequest) *constants.ErrorResponse
}

func NewLessonItemService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) LessonItemServiceInterface {
	return &service{
		repoCtx,
		infraCtx,
	}
}
