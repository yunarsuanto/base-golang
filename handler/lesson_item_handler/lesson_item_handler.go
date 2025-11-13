package lesson_item_handler

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

func (a handler) ListLessonItem(w http.ResponseWriter, r *http.Request) {
	var result listLessonItemResponse
	permission := constants.PermissionLessonItemList

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = listLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	var in listLessonItemRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeUrlQueryParams(&in, r.URL.Query())
	if errs != nil {
		result = listLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	pagination := objects.NewPagination()
	pagination.MapFromRequest(in.PaginationRequest)
	data, errs := a.LessonItemService.ListLessonItem(ctx, pagination, in.LessonId)
	if errs != nil {
		result = listLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := make([]*listLessonItemResponseData, len(data))
	for i, v := range data {
		resultData[i] = &listLessonItemResponseData{
			Id:       v.Id,
			LessonId: v.LessonId,
			Content:  v.Content,
			Order:    v.Order,
			Media:    v.Media,
			Group:    v.Group,
			IsDone:   v.IsDone,
		}
	}

	result = listLessonItemResponse{
		Meta:       utils.SetSuccessMeta("List LessonItem", permission),
		Pagination: pagination.MapToResponse(),
		Data:       resultData,
	}

	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) DetailLessonItem(w http.ResponseWriter, r *http.Request) {
	var result detailLessonItemResponse
	permission := constants.PermissionLessonItemDetail

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = detailLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	var in detailLessonItemRequest
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
		result = detailLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.DetailLessonItemRequest(in)
	data, errs := a.LessonItemService.DetailLessonItem(ctx, req)
	if errs != nil {
		result = detailLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := detailLessonItemResponseData{
		Id:      data.Id,
		Content: data.Content,
		Order:   data.Order,
		Media:   data.Media,
		Group:   data.Group,
		IsDone:  data.IsDone,
	}

	result = detailLessonItemResponse{
		Meta: utils.SetSuccessMeta("Detail Lesson Item", permission),
		Data: resultData,
	}

	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) CreateLessonItem(w http.ResponseWriter, r *http.Request) {
	var result createLessonItemResponse
	permission := constants.PermissionLessonItemCreate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = createLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in createLessonItemRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = createLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.CreateLessonItemRequest(in)
	errs = a.LessonItemService.CreateLessonItem(ctx, req)
	if errs != nil {
		result = createLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = createLessonItemResponse{
		Meta: utils.SetSuccessMeta("Create LessonItem", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
func (a handler) UpdateLessonItem(w http.ResponseWriter, r *http.Request) {
	var result updateLessonItemResponse
	permission := constants.PermissionLessonItemUpdate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = updateLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in updateLessonItemRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = updateLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.UpdateLessonItemRequest(in)
	errs = a.LessonItemService.UpdateLessonItem(ctx, req)
	if errs != nil {
		result = updateLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = updateLessonItemResponse{
		Meta: utils.SetSuccessMeta("Update LessonItem", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
func (a handler) DeleteLessonItem(w http.ResponseWriter, r *http.Request) {
	var result deleteLessonItemResponse
	permission := constants.PermissionLessonItemDelete

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = deleteLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in deleteLessonItemRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = deleteLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.DeleteLessonItemRequest(in)
	errs = a.LessonItemService.DeleteLessonItem(ctx, req)
	if errs != nil {
		result = deleteLessonItemResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = deleteLessonItemResponse{
		Meta: utils.SetSuccessMeta("Delete LessonItem", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
