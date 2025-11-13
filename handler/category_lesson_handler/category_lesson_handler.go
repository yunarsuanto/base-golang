package category_lesson_handler

import (
	"errors"
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

func (a handler) ListCategoryLessonPublic(w http.ResponseWriter, r *http.Request) {
	var result listCategoryLessonPublicResponse
	permission := constants.PermissionCategoryLessonList

	ctx := r.Context()

	var in listCategoryLessonRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs := utils.DecodeUrlQueryParams(&in, r.URL.Query())
	if errs != nil {
		result = listCategoryLessonPublicResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	pagination := objects.NewPagination()
	pagination.MapFromRequest(in.PaginationRequest)

	data, errs := a.CategoryLessonService.CategoryLessonPublic(ctx)
	if errs != nil {
		result = listCategoryLessonPublicResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	dataLessons, errs := a.LessonService.ListLesson(ctx, pagination, data.Id)
	if errs != nil {
		result = listCategoryLessonPublicResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	lessonIds := []string{}
	for _, v := range dataLessons {
		lessonIds = append(lessonIds, v.Id)
	}

	if len(lessonIds) == 0 {
		result = listCategoryLessonPublicResponse{Meta: utils.SetErrorMeta(&constants.ErrorResponse{HttpCode: 500, Err: errors.New("NOT DATA")}, permission)}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}

	dataLessonItems, errs := a.LessonItemService.ListLessonItemByLessonIds(ctx, lessonIds)
	if errs != nil {
		result = listCategoryLessonPublicResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultLessons := []listCategoryLessonPublicResponseDataLesson{}
	for _, v := range dataLessons {
		items := []listCategoryLessonPublicResponseDataLessonItem{}
		for _, item := range dataLessonItems {
			if item.LessonId == v.Id {
				items = append(items, listCategoryLessonPublicResponseDataLessonItem{
					Id:      item.Id,
					Content: item.Content,
					Order:   item.Order,
					Media:   item.Media,
				})
			}
		}
		resultLessons = append(resultLessons, listCategoryLessonPublicResponseDataLesson{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Level:       v.Level,
			Media:       v.Media,
			Items:       items,
		})
	}

	resultData := listCategoryLessonPublicResponseData{
		Id:                 data.Id,
		Title:              data.Title,
		Description:        data.Description,
		CategoryLessonType: data.CategoryLessonType,
		Media:              data.Media,
		Lessons:            resultLessons,
	}

	result = listCategoryLessonPublicResponse{
		Meta: utils.SetSuccessMeta("List Category Lesson Public", permission),
		Data: resultData,
	}

	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) ListCategoryLesson(w http.ResponseWriter, r *http.Request) {
	var result listCategoryLessonResponse
	permission := constants.PermissionCategoryLessonList

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = listCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	var in listCategoryLessonRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeUrlQueryParams(&in, r.URL.Query())
	if errs != nil {
		result = listCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	pagination := objects.NewPagination()
	pagination.MapFromRequest(in.PaginationRequest)
	data, errs := a.CategoryLessonService.ListCategoryLesson(ctx, pagination)
	if errs != nil {
		result = listCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*listCategoryLessonResponseData{}
	for _, v := range data {
		resultData = append(resultData, &listCategoryLessonResponseData{
			Id:                 v.Id,
			Title:              v.Title,
			Description:        v.Description,
			CategoryLessonType: v.CategoryLessonType,
			Media:              v.Media,
		})
	}

	result = listCategoryLessonResponse{
		Meta:       utils.SetSuccessMeta("List Category Lesson", permission),
		Pagination: pagination.MapToResponse(),
		Data:       resultData,
	}

	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) DetailCategoryLesson(w http.ResponseWriter, r *http.Request) {
	var result detailCategoryLessonResponse
	permission := constants.PermissionCategoryLessonDetail

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = detailCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	var in detailCategoryLessonRequest
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
		result = detailCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.DetailCategoryLessonRequest(in)
	data, errs := a.CategoryLessonService.DetailCategoryLesson(ctx, req)
	if errs != nil {
		result = detailCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := detailCategoryLessonDataResponse{
		Id:                 data.Id,
		Title:              data.Title,
		Description:        data.Description,
		CategoryLessonType: data.CategoryLessonType,
		Media:              data.Media,
	}

	result = detailCategoryLessonResponse{
		Meta: utils.SetSuccessMeta("Detail Category Lesson", permission),
		Data: resultData,
	}

	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) CreateCategoryLesson(w http.ResponseWriter, r *http.Request) {
	var result createCategoryLessonResponse
	permission := constants.PermissionCategoryLessonCreate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = createCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in createCategoryLessonRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = createCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.CreateCategoryLessonRequest(in)
	errs = a.CategoryLessonService.CreateCategoryLesson(ctx, req)
	if errs != nil {
		result = createCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = createCategoryLessonResponse{
		Meta: utils.SetSuccessMeta("Create Category Lesson", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
func (a handler) UpdateCategoryLesson(w http.ResponseWriter, r *http.Request) {
	var result updateCategoryLessonResponse
	permission := constants.PermissionCategoryLessonUpdate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = updateCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in updateCategoryLessonRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = updateCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.UpdateCategoryLessonRequest(in)
	errs = a.CategoryLessonService.UpdateCategoryLesson(ctx, req)
	if errs != nil {
		result = updateCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = updateCategoryLessonResponse{
		Meta: utils.SetSuccessMeta("Update Category Lesson", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
func (a handler) DeleteCategoryLesson(w http.ResponseWriter, r *http.Request) {
	var result deleteCategoryLessonResponse
	permission := constants.PermissionCategoryLessonDelete

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = deleteCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in deleteCategoryLessonRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = deleteCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.DeleteCategoryLessonRequest(in)
	errs = a.CategoryLessonService.DeleteCategoryLesson(ctx, req)
	if errs != nil {
		result = deleteCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = deleteCategoryLessonResponse{
		Meta: utils.SetSuccessMeta("Delete Category Lesson", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
