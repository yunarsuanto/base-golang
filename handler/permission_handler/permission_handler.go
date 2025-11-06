package permission_handler

import (
	"net/http"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/service"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type handler struct {
	*service.ServiceCtx
}

func (a handler) ListPermission(w http.ResponseWriter, r *http.Request) {
	var result listPermissionResponse
	permission := constants.PermissionPermissionList

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = listPermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	var in listPermissionRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeUrlQueryParams(&in, r.URL.Query())
	if errs != nil {
		result = listPermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	pagination := objects.NewPagination()
	pagination.MapFromRequest(in.PaginationRequest)
	data, errs := a.PermissionService.ListPermission(ctx, pagination)
	if errs != nil {
		result = listPermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := make([]*listPermissionResponseData, len(data))
	for i, v := range data {
		resultData[i] = &listPermissionResponseData{
			Id:   v.Id,
			Name: v.Name,
		}
	}

	result = listPermissionResponse{
		Meta:       utils.SetSuccessMeta("List Permission", permission),
		Pagination: pagination.MapToResponse(),
		Data:       resultData,
	}

	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) CreatePermission(w http.ResponseWriter, r *http.Request) {
	var result createPermissionResponse
	permission := constants.PermissionPermissionCreate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = createPermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in createPermissionRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = createPermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.CreatePermissionRequest(in)
	errs = a.PermissionService.CreatePermission(ctx, req)
	if errs != nil {
		result = createPermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = createPermissionResponse{
		Meta: utils.SetSuccessMeta("Create Permission", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
func (a handler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	var result updatePermissionResponse
	permission := constants.PermissionPermissionUpdate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = updatePermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in updatePermissionRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = updatePermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.UpdatePermissionRequest(in)
	errs = a.PermissionService.UpdatePermission(ctx, req)
	if errs != nil {
		result = updatePermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = updatePermissionResponse{
		Meta: utils.SetSuccessMeta("Update Permission", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
func (a handler) DeletePermission(w http.ResponseWriter, r *http.Request) {
	var result deletePermissionResponse
	permission := constants.PermissionPermissionDelete

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = deletePermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in deletePermissionRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = deletePermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.DeletePermissionRequest(in)
	errs = a.PermissionService.DeletePermission(ctx, req)
	if errs != nil {
		result = deletePermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = deletePermissionResponse{
		Meta: utils.SetSuccessMeta("Delete Permission", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
