package role_permission_handler

import (
	"github.com/yunarsuanto/base-go/constants"
	common_input_handler "github.com/yunarsuanto/base-go/handler"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
	"golang.org/x/net/context"
)

type upsertRolePermissionRequest struct {
	RoleId       string `json:"role_id" schema:"role_id" validate:"required,uuid"`
	PermissionId string `json:"permission_id" schema:"permission_id" validate:"required,uuid"`
}

type upsertRolePermissionResponse struct {
	Meta common_input_handler.Meta         `json:"meta"`
	Data *upsertRolePermissionResponseData `json:"data"`
}

type upsertRolePermissionResponseData struct {
}

type deleteRolePermissionRequest struct {
	RoleId       string `json:"role_id" schema:"role_id" validate:"required,uuid"`
	PermissionId string `json:"permission_id" schema:"permission_id" validate:"required,uuid"`
}

type deleteRolePermissionResponse struct {
	Meta common_input_handler.Meta         `json:"meta"`
	Data *deleteRolePermissionResponseData `json:"data"`
}

type deleteRolePermissionResponseData struct {
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
