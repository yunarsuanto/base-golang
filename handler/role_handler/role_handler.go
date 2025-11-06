package role_handler

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

func (a handler) ListRole(w http.ResponseWriter, r *http.Request) {
	var result listRoleResponse
	permission := constants.PermissionRoleList

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = listRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	var in listRoleRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeUrlQueryParams(&in, r.URL.Query())
	if errs != nil {
		result = listRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	pagination := objects.NewPagination()
	pagination.MapFromRequest(in.PaginationRequest)
	data, errs := a.RoleService.ListRole(ctx, pagination)
	if errs != nil {
		result = listRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := make([]*listRoleResponseData, len(data))
	for i, v := range data {
		resultData[i] = &listRoleResponseData{
			Id:   v.Id,
			Name: v.Name,
		}
	}

	result = listRoleResponse{
		Meta:       utils.SetSuccessMeta("List Role", permission),
		Pagination: pagination.MapToResponse(),
		Data:       resultData,
	}

	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) DetailRole(w http.ResponseWriter, r *http.Request) {
	var result detailRoleResponse
	permission := constants.PermissionRoleDetail

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = detailRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	var in detailRoleRequest
	vars := mux.Vars(r)
	in.Id = vars["id"]
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	req := objects.DetailRoleRequest(in)
	data, errs := a.RoleService.DetailRole(ctx, req)
	if errs != nil {
		result = detailRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := detailRoleResponseData{}
	permissions := []detailRoleResponseDataPermission{}
	for _, v := range data {
		resultData.Id = v.Id
		resultData.Name = v.Name
		if utils.NullScan(v.PermissionId) != "" {
			permissions = append(permissions, detailRoleResponseDataPermission{
				Id:   utils.NullScan(v.PermissionId),
				Name: utils.NullScan(v.PermissionName),
			})
		}
		resultData.Permissions = &permissions
	}

	result = detailRoleResponse{
		Meta: utils.SetSuccessMeta("Detail Role", permission),
		Data: &resultData,
	}

	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var result createRoleResponse
	permission := constants.PermissionRoleCreate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = createRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in createRoleRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = createRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.CreateRoleRequest(in)
	errs = a.RoleService.CreateRole(ctx, req)
	if errs != nil {
		result = createRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = createRoleResponse{
		Meta: utils.SetSuccessMeta("Create Role", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
func (a handler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	var result updateRoleResponse
	permission := constants.PermissionRoleUpdate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = updateRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in updateRoleRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = updateRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.UpdateRoleRequest(in)
	errs = a.RoleService.UpdateRole(ctx, req)
	if errs != nil {
		result = updateRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = updateRoleResponse{
		Meta: utils.SetSuccessMeta("Update Role", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
func (a handler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	var result deleteRoleResponse
	permission := constants.PermissionRoleDelete

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = deleteRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in deleteRoleRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = deleteRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.DeleteRoleRequest(in)
	errs = a.RoleService.DeleteRole(ctx, req)
	if errs != nil {
		result = deleteRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = deleteRoleResponse{
		Meta: utils.SetSuccessMeta("Delete Role", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
