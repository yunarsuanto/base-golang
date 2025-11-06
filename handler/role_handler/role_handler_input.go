package role_handler

import (
	"github.com/yunarsuanto/base-go/constants"
	common_input_handler "github.com/yunarsuanto/base-go/handler"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
	"golang.org/x/net/context"
)

type listRoleRequest struct {
	common_input_handler.PaginationRequest
}

type listRoleResponse struct {
	Meta       common_input_handler.Meta        `json:"meta"`
	Pagination *common_input_handler.Pagination `json:"pagination"`
	Data       []*listRoleResponseData          `json:"data"`
}

type listRoleResponseData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type createRoleRequest struct {
	Name string `schema:"name" validate:"required"`
}

type createRoleResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *createRoleResponseData   `json:"data"`
}

type createRoleResponseData struct {
}

type updateRoleRequest struct {
	Id   string `schema:"id" validate:"required,uuid"`
	Name string `schema:"name" validate:"required"`
}

type updateRoleResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *updateRoleResponseData   `json:"data"`
}

type updateRoleResponseData struct {
}

type deleteRoleRequest struct {
	Id string `schema:"id" validate:"required,uuid"`
}

type deleteRoleResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *deleteRoleResponseData   `json:"data"`
}

type deleteRoleResponseData struct {
}

type detailRoleRequest struct {
	Id string `schema:"id" validate:"required,uuid"`
}

type detailRoleResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *detailRoleResponseData   `json:"data"`
}

type detailRoleResponseData struct {
	Id          string                              `json:"id"`
	Name        string                              `json:"name"`
	Permissions *[]detailRoleResponseDataPermission `json:"permissions"`
}

type detailRoleResponseDataPermission struct {
	Id   string `json:"id"`
	Name string `json:"name"`
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
