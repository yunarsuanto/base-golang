package lesson_service

import (
	"context"
	"fmt"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type service struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (a service) ListLesson(ctx context.Context, pagination *objects.Pagination, categoryLessonId string) ([]objects.ListLessonResponse, *constants.ErrorResponse) {
	var result []objects.ListLessonResponse

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	resultData, errs := a.LessonRepo.ListLesson(ctx, tx, pagination, categoryLessonId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	for _, v := range resultData {
		result = append(result, objects.ListLessonResponse(v))
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a service) DetailLesson(ctx context.Context, req objects.DetailLessonRequest) (objects.DetailLessonResponse, *constants.ErrorResponse) {
	var result objects.DetailLessonResponse

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	resultData, errs := a.LessonRepo.DetailLesson(ctx, tx, req.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.DetailLessonResponse(resultData)
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a service) CreateLesson(ctx context.Context, req objects.CreateLessonRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	createData := models.CreateLesson{
		Title:            req.Title,
		Description:      req.Description,
		CategoryLessonId: req.CategoryLessonId,
		LessonType:       req.LessonType,
		Media:            req.Media,
		Level:            req.Level,
	}

	_, errs := a.LessonRepo.CreateLesson(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return utils.ErrorInternalServer(err.Error())
	}

	return nil
}
func (a service) UpdateLesson(ctx context.Context, req objects.UpdateLessonRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	updateData := models.UpdateLesson{
		Id:               req.Id,
		Title:            req.Title,
		Description:      req.Description,
		CategoryLessonId: req.CategoryLessonId,
		LessonType:       req.LessonType,
		Media:            req.Media,
		Level:            req.Level,
	}

	errs := a.LessonRepo.UpdateLesson(ctx, tx, updateData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return utils.ErrorInternalServer(err.Error())
	}

	return nil
}
func (a service) DeleteLesson(ctx context.Context, req objects.DeleteLessonRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	deleteData := models.DeleteLesson{
		Id: req.Id,
	}

	errs := a.LessonRepo.DeleteLesson(ctx, tx, deleteData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return utils.ErrorInternalServer(err.Error())
	}

	return nil
}

func (a service) CopyLessonItem(ctx context.Context, req objects.CopyLessonRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	var lessonId string
	var dataItem []models.ListLessonItem

	detail, errs := a.LessonRepo.DetailLesson(ctx, tx, req.LessonId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	lessonId, errs = a.LessonRepo.CreateLesson(ctx, tx, models.CreateLesson{
		Title:            detail.Title + fmt.Sprintf(" Level %d", req.Level),
		Description:      detail.Description,
		CategoryLessonId: detail.CategoryLessonId,
		LessonType:       detail.LessonType,
		Media:            detail.Media,
		Level:            req.Level,
	})
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	dataItem, errs = a.LessonItemRepo.ListLessonItem(ctx, tx, &objects.Pagination{
		Page:  1,
		Limit: 1000,
	}, req.LessonId)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	if len(dataItem) == 0 {
		_ = tx.Rollback()
		return errs
	}

	dataInsert := []models.BulkCreateLessonItem{}

	for _, v := range dataItem {
		dataInsert = append(dataInsert, models.BulkCreateLessonItem{
			LessonId: lessonId,
			Content:  v.Content,
			Order:    v.Order,
			Media:    v.Media,
			Group:    v.Group,
			IsDone:   v.IsDone,
		})
	}

	errs = a.LessonItemRepo.BulkCreateLessonItem(ctx, tx, dataInsert)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return utils.ErrorInternalServer(err.Error())
	}

	return nil
}
