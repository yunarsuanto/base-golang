package lesson_item_service

import (
	"context"

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

func (a service) ListLessonItem(ctx context.Context, pagination *objects.Pagination, lessonId string) ([]objects.ListLessonItemResponse, *constants.ErrorResponse) {
	var result []objects.ListLessonItemResponse

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	resultData, errs := a.LessonItemRepo.ListLessonItem(ctx, tx, pagination, lessonId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	for _, v := range resultData {
		result = append(result, objects.ListLessonItemResponse(v))
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a service) ListLessonItemByLessonIds(ctx context.Context, lessonIds []string) ([]objects.ListLessonItemResponse, *constants.ErrorResponse) {
	var result []objects.ListLessonItemResponse

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	resultData, errs := a.LessonItemRepo.ListLessonItemByLessonIds(ctx, tx, lessonIds)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	for _, v := range resultData {
		result = append(result, objects.ListLessonItemResponse(v))
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a service) DetailLessonItem(ctx context.Context, req objects.DetailLessonItemRequest) (objects.DetailLessonItemResponse, *constants.ErrorResponse) {
	var result objects.DetailLessonItemResponse

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	resultData, errs := a.LessonItemRepo.DetailLessonItem(ctx, tx, req.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.DetailLessonItemResponse(resultData)

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a service) CreateLessonItem(ctx context.Context, req objects.CreateLessonItemRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	createData := models.CreateLessonItem{
		LessonId: req.LessonId,
		Content:  req.Content,
		Order:    req.Order,
		Media:    req.Media,
		Group:    req.Group,
		IsDone:   req.IsDone,
	}

	errs := a.LessonItemRepo.CreateLessonItem(ctx, tx, createData)
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
func (a service) UpdateLessonItem(ctx context.Context, req objects.UpdateLessonItemRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	updateData := models.UpdateLessonItem{
		Id:       req.Id,
		LessonId: req.LessonId,
		Content:  req.Content,
		Order:    req.Order,
		Media:    req.Media,
		Group:    req.Group,
		IsDone:   req.IsDone,
	}

	errs := a.LessonItemRepo.UpdateLessonItem(ctx, tx, updateData)
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
func (a service) DeleteLessonItem(ctx context.Context, req objects.DeleteLessonItemRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	deleteData := models.DeleteLessonItem{
		Id: req.Id,
	}

	errs := a.LessonItemRepo.DeleteLessonItem(ctx, tx, deleteData)
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
