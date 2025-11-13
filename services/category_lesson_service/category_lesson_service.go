package category_lesson_service

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

func (a service) ListCategoryLesson(ctx context.Context, pagination *objects.Pagination) ([]objects.ListCategoryLessonResponse, *constants.ErrorResponse) {
	var result []objects.ListCategoryLessonResponse

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	resultData, errs := a.CategoryLessonRepo.ListCategoryLesson(ctx, tx, pagination)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	for _, v := range resultData {
		result = append(result, objects.ListCategoryLessonResponse(v))
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a service) DetailCategoryLesson(ctx context.Context, req objects.DetailCategoryLessonRequest) (objects.DetailCategoryLessonResponse, *constants.ErrorResponse) {
	var result objects.DetailCategoryLessonResponse

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	resultData, errs := a.CategoryLessonRepo.DetailCategoryLesson(ctx, tx, req.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.DetailCategoryLessonResponse(resultData)

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a service) CreateCategoryLesson(ctx context.Context, req objects.CreateCategoryLessonRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	createData := models.CreateCategoryLesson{
		Title:              req.Title,
		Description:        req.Description,
		CategoryLessonType: req.CategoryLessonType,
		Media:              req.Media,
	}

	errs := a.CategoryLessonRepo.CreateCategoryLesson(ctx, tx, createData)
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
func (a service) UpdateCategoryLesson(ctx context.Context, req objects.UpdateCategoryLessonRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	updateData := models.UpdateCategoryLesson{
		Id:                 req.Id,
		Title:              req.Title,
		Description:        req.Description,
		CategoryLessonType: req.CategoryLessonType,
		Media:              req.Media,
	}

	errs := a.CategoryLessonRepo.UpdateCategoryLesson(ctx, tx, updateData)
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
func (a service) DeleteCategoryLesson(ctx context.Context, req objects.DeleteCategoryLessonRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	deleteData := models.DeleteCategoryLesson{
		Id: req.Id,
	}

	errs := a.CategoryLessonRepo.DeleteCategoryLesson(ctx, tx, deleteData)
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

func (a service) CategoryLessonPublic(ctx context.Context) (objects.ListCategoryLessonPublicResponse, *constants.ErrorResponse) {
	var result objects.ListCategoryLessonPublicResponse

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	resultData, errs := a.CategoryLessonRepo.CategoryLessonPublic(ctx, tx)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.ListCategoryLessonPublicResponse(resultData)

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}
