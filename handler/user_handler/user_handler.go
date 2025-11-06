package user_handler

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

func (a handler) ListUser(w http.ResponseWriter, r *http.Request) {
	var result listUserResponse
	permission := constants.PermissionUserList

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = listUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	var in listUserRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeUrlQueryParams(&in, r.URL.Query())
	if errs != nil {
		result = listUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	pagination := objects.NewPagination()
	pagination.MapFromRequest(in.PaginationRequest)
	data, errs := a.UserService.ListUser(ctx, pagination)
	if errs != nil {
		result = listUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := make([]*listUserResponseData, len(data))
	for i, v := range data {
		resultData[i] = &listUserResponseData{
			Id:       v.Id,
			Username: v.Username,
		}
	}

	result = listUserResponse{
		Meta:       utils.SetSuccessMeta("List User", permission),
		Pagination: pagination.MapToResponse(),
		Data:       resultData,
	}

	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) DetailUser(w http.ResponseWriter, r *http.Request) {
	var result detailUserResponse
	permission := constants.PermissionUserDetail

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = detailUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	var in detailUserRequest
	vars := mux.Vars(r)
	in.Id = vars["id"]
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	req := objects.DetailUserRequest(in)
	data, errs := a.UserService.DetailUser(ctx, req)
	if errs != nil {
		result = detailUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	rolesMap := make(map[string]detailUserResponseDataRole)
	resultData := detailUserResponseData{}

	for _, v := range data {
		resultData.Id = v.Id
		resultData.Username = v.Username
		resultData.IsActive = v.IsActive
		roleId := utils.NullScan(v.RoleId)
		role, exists := rolesMap[roleId]
		if !exists {
			role = detailUserResponseDataRole{
				Id:          roleId,
				Name:        utils.NullScan(v.RoleName),
				IsActive:    utils.NullScan(v.RoleIsActive),
				Permissions: &[]detailUserResponseDataRolePermission{},
			}
		}
		rolesMap[roleId] = role

		permissions := *role.Permissions
		if utils.NullScan(v.PermissionId) != "" {
			permissions = append(permissions, detailUserResponseDataRolePermission{
				Id:   utils.NullScan(v.PermissionId),
				Name: utils.NullScan(v.PermissionName),
			})
		}
		role.Permissions = &permissions
		rolesMap[roleId] = role
	}

	roles := []detailUserResponseDataRole{}
	for _, role := range rolesMap {
		if role.Id != "" {
			roles = append(roles, detailUserResponseDataRole{
				Id:          role.Id,
				Name:        role.Name,
				IsActive:    role.IsActive,
				Permissions: role.Permissions,
			})
		}
	}

	resultData.Roles = &roles

	result = detailUserResponse{
		Meta: utils.SetSuccessMeta("Detail User", permission),
		Data: &resultData,
	}

	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var result createUserResponse
	permission := constants.PermissionUserCreate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = createUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in createUserRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = createUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.CreateUserRequest(in)
	errs = a.UserService.CreateUser(ctx, req)
	if errs != nil {
		result = createUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = createUserResponse{
		Meta: utils.SetSuccessMeta("Create User", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
func (a handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var result updateUserResponse
	permission := constants.PermissionUserUpdate

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = updateUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in updateUserRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = updateUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.UpdateUserRequest(in)
	errs = a.UserService.UpdateUser(ctx, req)
	if errs != nil {
		result = updateUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = updateUserResponse{
		Meta: utils.SetSuccessMeta("Update User", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
func (a handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var result deleteUserResponse
	permission := constants.PermissionUserDelete

	ctx := r.Context()
	errs := a.checkPermission(ctx, permission)
	if errs != nil {
		result = deleteUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	var in deleteUserRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs = utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = deleteUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	req := objects.DeleteUserRequest(in)
	errs = a.UserService.DeleteUser(ctx, req)
	if errs != nil {
		result = deleteUserResponse{Meta: utils.SetErrorMeta(errs, permission)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = deleteUserResponse{
		Meta: utils.SetSuccessMeta("Delete User", permission),
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
