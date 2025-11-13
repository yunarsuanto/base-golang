package lesson_service

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
)

type LessonServiceInterface interface {
	ListLesson(ctx context.Context, pagination *objects.Pagination, categoryLessonId string) ([]objects.ListLessonResponse, *constants.ErrorResponse)
	DetailLesson(ctx context.Context, req objects.DetailLessonRequest) (objects.DetailLessonResponse, *constants.ErrorResponse)
	CreateLesson(ctx context.Context, req objects.CreateLessonRequest) *constants.ErrorResponse
	UpdateLesson(ctx context.Context, req objects.UpdateLessonRequest) *constants.ErrorResponse
	DeleteLesson(ctx context.Context, req objects.DeleteLessonRequest) *constants.ErrorResponse
	CopyLessonItem(ctx context.Context, req objects.CopyLessonRequest) *constants.ErrorResponse
}

func NewLessonService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) LessonServiceInterface {
	return &service{
		repoCtx,
		infraCtx,
	}
}
