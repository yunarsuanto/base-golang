package permission_handler

import (
	"github.com/yunarsuanto/base-go/constants"
	common_input_handler "github.com/yunarsuanto/base-go/handler"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
	"golang.org/x/net/context"
)

type listPermissionRequest struct {
	common_input_handler.PaginationRequest
}

type listPermissionResponse struct {
	Meta       common_input_handler.Meta        `json:"meta"`
	Pagination *common_input_handler.Pagination `json:"pagination"`
	Data       []*listPermissionResponseData    `json:"data"`
}

type listPermissionResponseData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type createPermissionRequest struct {
	Name string `schema:"name" validate:"required"`
}

type createPermissionResponse struct {
	Meta common_input_handler.Meta     `json:"meta"`
	Data *createPermissionResponseData `json:"data"`
}

type createPermissionResponseData struct {
}

type updatePermissionRequest struct {
	Id   string `schema:"id" validate:"required,uuid"`
	Name string `schema:"name" validate:"required"`
}

type updatePermissionResponse struct {
	Meta common_input_handler.Meta     `json:"meta"`
	Data *updatePermissionResponseData `json:"data"`
}

type updatePermissionResponseData struct {
}

type deletePermissionRequest struct {
	Id string `schema:"id" validate:"required,uuid"`
}

type deletePermissionResponse struct {
	Meta common_input_handler.Meta     `json:"meta"`
	Data *deletePermissionResponseData `json:"data"`
}

type deletePermissionResponseData struct {
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
