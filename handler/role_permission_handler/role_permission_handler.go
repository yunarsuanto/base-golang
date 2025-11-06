package role_permission_handler

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

func (a handler) UpsertRolePermission(w http.ResponseWriter, r *http.Request) {
	var result upsertRolePermissionResponse
	permission := constants.PermissionRolePermissionUpsert

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = upsertRolePermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in upsertRolePermissionRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = upsertRolePermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.UpsertRolePermissionRequest(in)
	errs = a.RolePermissionService.UpsertRolePermission(ctx, req)
	if errs != nil {
		result = upsertRolePermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = upsertRolePermissionResponse{
		Meta: utils.SetSuccessMeta("Upsert RolePermission", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) DeleteRolePermission(w http.ResponseWriter, r *http.Request) {
	var result deleteRolePermissionResponse
	permission := constants.PermissionRolePermissionDelete

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = deleteRolePermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in deleteRolePermissionRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = deleteRolePermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.DeleteRolePermissionRequest(in)
	errs = a.RolePermissionService.DeleteRolePermission(ctx, req)
	if errs != nil {
		result = deleteRolePermissionResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = deleteRolePermissionResponse{
		Meta: utils.SetSuccessMeta("Delete RolePermission", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
