package category_lesson_handler

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
	data, errs := a.CategoryLessonService.ListCategoryLesson(ctx, pagination, in.HasParent)
	if errs != nil {
		result = listCategoryLessonResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*listCategoryLessonResponseData{}
	for _, v := range data {
		resultData = append(resultData, &listCategoryLessonResponseData{
			Id:               v.Id,
			Title:            v.Title,
			Description:      v.Description,
			Media:            v.Media,
			CategoryLessonId: utils.NullScan(v.CategoryLessonId),
		})
	}
	// childMap := map[string][]listCategoryLessonResponseData{}
	// for _, v := range data {
	// 	if parentId := utils.NullScan(v.CategoryLessonId); parentId != "" {
	// 		childMap[parentId] = append(childMap[parentId], listCategoryLessonResponseData{
	// 			Id:               v.Id,
	// 			Title:            v.Title,
	// 			Description:      v.Description,
	// 			Media:            v.Media,
	// 			CategoryLessonId: parentId,
	// 		})
	// 	}
	// }
	// for _, v := range data {
	// 	if utils.NullScan(v.CategoryLessonId) == "" {
	// 		resultData = append(resultData, &listCategoryLessonResponseData{
	// 			Id:                  v.Id,
	// 			Title:               v.Title,
	// 			Description:         v.Description,
	// 			Media:               v.Media,
	// 			CategoryLessonChild: childMap[v.Id],
	// 		})
	// 	}
	// }

	result = listCategoryLessonResponse{
		Meta:       utils.SetSuccessMeta("List CategoryLesson", permission),
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

	resultData := detailCategoryLessonDataResponse{}
	childs := []listCategoryLessonResponseData{}
	for _, v := range data {
		if req.Id == v.Id {
			resultData = detailCategoryLessonDataResponse{
				Id:               v.Id,
				Title:            v.Title,
				Description:      v.Description,
				Media:            v.Media,
				CategoryLessonId: utils.NullScan(v.CategoryLessonId),
			}
		} else if utils.NullScan(v.CategoryLessonId) == req.Id {
			childs = append(childs, listCategoryLessonResponseData{
				Id:               v.Id,
				Title:            v.Title,
				Description:      v.Description,
				Media:            v.Media,
				CategoryLessonId: utils.NullScan(v.CategoryLessonId),
			})
		}
	}
	resultData.Childs = childs
	result = detailCategoryLessonResponse{
		Meta: utils.SetSuccessMeta("Detail CategoryLesson", permission),
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
		Meta: utils.SetSuccessMeta("Create CategoryLesson", permission),
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
		Meta: utils.SetSuccessMeta("Update CategoryLesson", permission),
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
		Meta: utils.SetSuccessMeta("Delete CategoryLesson", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
