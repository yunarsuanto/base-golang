package user_role_handler

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

func (a handler) UpsertUserRole(w http.ResponseWriter, r *http.Request) {
	var result upsertUserRoleResponse
	permission := constants.PermissionUserRoleUpsert

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = upsertUserRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in upsertUserRoleRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = upsertUserRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.UpsertUserRoleRequest(in)
	errs = a.UserRoleService.UpdateUserRole(ctx, req)
	if errs != nil {
		result = upsertUserRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = upsertUserRoleResponse{
		Meta: utils.SetSuccessMeta("Upsert UserRole", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) DeleteUserRole(w http.ResponseWriter, r *http.Request) {
	var result deleteUserRoleResponse
	permission := constants.PermissionUserRoleDelete

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = deleteUserRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in deleteUserRoleRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = deleteUserRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.DeleteUserRoleRequest(in)
	errs = a.UserRoleService.DeleteUserRole(ctx, req)
	if errs != nil {
		result = deleteUserRoleResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = deleteUserRoleResponse{
		Meta: utils.SetSuccessMeta("Delete UserRole", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
