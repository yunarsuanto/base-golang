package lesson_handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/service"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type handler struct {
	*service.ServiceCtx
}

func (a handler) ListLesson(w http.ResponseWriter, r *http.Request) {
	var result listLessonResponse
	permission := constants.PermissionLessonList

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = listLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	var in listLessonRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeUrlQueryParams(&in, r.URL.Query())
	if errs != nil {
		result = listLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	pagination := objects.NewPagination()
	pagination.MapFromRequest(in.PaginationRequest)
	data, errs := a.LessonService.ListLesson(ctx, pagination, in.CategoryLessonId)
	if errs != nil {
		result = listLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := make([]*listLessonResponseData, len(data))
	for i, v := range data {
		resultData[i] = &listLessonResponseData{
			Id:                  v.Id,
			CategoryLessonId:    v.CategoryLessonId,
			Title:               v.Title,
			Description:         v.Description,
			Media:               v.Media,
			CategoryLessonTitle: v.CategoryLessonTitle,
			LessonType:          v.LessonType,
			Level:               v.Level,
		}
	}

	result = listLessonResponse{
		Meta:       utils.SetSuccessMeta("List Lesson", permission),
		Pagination: pagination.MapToResponse(),
		Data:       resultData,
	}

	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) DetailLesson(w http.ResponseWriter, r *http.Request) {
	var result detailLessonResponse
	permission := constants.PermissionLessonDetail

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = detailLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	var in detailLessonRequest
	vars := mux.Vars(r)
	in.Id = vars["id"]
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeUrlQueryParams(&in, r.URL.Query())
	if errs != nil {
		result = detailLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.DetailLessonRequest(in)
	data, errs := a.LessonService.DetailLesson(ctx, req)
	if errs != nil {
		result = detailLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := detailLessonResponseData{
		Id:               data.Id,
		Title:            data.Title,
		Description:      data.Description,
		CategoryLessonId: data.CategoryLessonId,
		LessonType:       data.LessonType,
		Media:            data.Media,
		Level:            data.Level,
	}
	result = detailLessonResponse{
		Meta: utils.SetSuccessMeta("Detail Lesson", permission),
		Data: resultData,
	}

	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) CreateLesson(w http.ResponseWriter, r *http.Request) {
	var result createLessonResponse
	permission := constants.PermissionLessonCreate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = createLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in createLessonRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = createLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.CreateLessonRequest(in)
	errs = a.LessonService.CreateLesson(ctx, req)
	if errs != nil {
		result = createLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = createLessonResponse{
		Meta: utils.SetSuccessMeta("Create Lesson", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
func (a handler) UpdateLesson(w http.ResponseWriter, r *http.Request) {
	var result updateLessonResponse
	permission := constants.PermissionLessonUpdate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = updateLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in updateLessonRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = updateLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.UpdateLessonRequest(in)
	errs = a.LessonService.UpdateLesson(ctx, req)
	if errs != nil {
		result = updateLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = updateLessonResponse{
		Meta: utils.SetSuccessMeta("Update Lesson", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
func (a handler) DeleteLesson(w http.ResponseWriter, r *http.Request) {
	var result deleteLessonResponse
	permission := constants.PermissionLessonDelete

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = deleteLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in deleteLessonRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = deleteLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.DeleteLessonRequest(in)
	errs = a.LessonService.DeleteLesson(ctx, req)
	if errs != nil {
		result = deleteLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = deleteLessonResponse{
		Meta: utils.SetSuccessMeta("Delete Lesson", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) CopyLessonItem(w http.ResponseWriter, r *http.Request) {
	var result copyLessonResponse
	permission := constants.PermissionLessonUpdate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = copyLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in copyLessonRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = copyLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.CopyLessonRequest(in)
	errs = a.LessonService.CopyLessonItem(ctx, req)
	if errs != nil {
		result = copyLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = copyLessonResponse{
		Meta: utils.SetSuccessMeta("Copy Lesson", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
