package user_role_handler

import (
	"github.com/yunarsuanto/base-go/constants"
	common_input_handler "github.com/yunarsuanto/base-go/handler"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
	"golang.org/x/net/context"
)

type upsertUserRoleRequest struct {
	UserId   string `json:"user_id" schema:"user_id" validate:"required,uuid"`
	RoleId   string `json:"role_id" schema:"role_id" validate:"required"`
	IsActive bool   `json:"is_active" schema:"is_active" validate:"required"`
}

type upsertUserRoleResponse struct {
	Meta common_input_handler.Meta   `json:"meta"`
	Data *upsertUserRoleResponseData `json:"data"`
}

type upsertUserRoleResponseData struct {
}

type deleteUserRoleRequest struct {
	UserId string `json:"user_id" schema:"user_id" validate:"required,uuid"`
	RoleId string `json:"role_id" schema:"role_id" validate:"required"`
}

type deleteUserRoleResponse struct {
	Meta common_input_handler.Meta   `json:"meta"`
	Data *deleteUserRoleResponseData `json:"data"`
}

type deleteUserRoleResponseData struct {
}

func (a handler) checkPermission(ctx context.Context, permission string) *constants.ErrorResponse {
	claims, ok := ctx.Value(constants.ClaimsContextKey).(*objects.JWTClaims)
	if !ok || claims == nil {
		return constants.ErrTokenInvalid
	}

	if !claims.IsSuperAdmin {
		if !utils.InArrayExist(permission, claims.Permissions) {
			return constants.ErrIneligibleAccess
		}
	}

	return nil
}
