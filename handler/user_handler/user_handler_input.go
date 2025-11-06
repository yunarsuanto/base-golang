package user_handler

import (
	"fmt"

	"github.com/yunarsuanto/base-go/constants"
	common_input_handler "github.com/yunarsuanto/base-go/handler"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
	"golang.org/x/net/context"
)

type listUserRequest struct {
	common_input_handler.PaginationRequest
}

type listUserResponse struct {
	Meta       common_input_handler.Meta        `json:"meta"`
	Pagination *common_input_handler.Pagination `json:"pagination"`
	Data       []*listUserResponseData          `json:"data"`
}

type listUserResponseData struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type detailUserRequest struct {
	Id string `schema:"id" validate:"required,uuid"`
}

type detailUserResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *detailUserResponseData   `json:"data"`
}

type detailUserResponseData struct {
	Id       string                        `json:"id"`
	Username string                        `json:"username"`
	IsActive bool                          `json:"is_active"`
	Roles    *[]detailUserResponseDataRole `json:"roles"`
}

type detailUserResponseDataRole struct {
	Id          string                                  `json:"id"`
	Name        string                                  `json:"name"`
	IsActive    bool                                    `json:"is_active"`
	Permissions *[]detailUserResponseDataRolePermission `json:"permissions"`
}

type detailUserResponseDataRolePermission struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type createUserRequest struct {
	Username string `schema:"username" validate:"required"`
	Password string `schema:"password" validate:"required"`
}

type createUserResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *createUserResponseData   `json:"data"`
}

type createUserResponseData struct {
}

type updateUserRequest struct {
	Id       string `schema:"id" validate:"required,uuid"`
	Username string `schema:"username" validate:"required"`
}

type updateUserResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *updateUserResponseData   `json:"data"`
}

type updateUserResponseData struct {
}

type deleteUserRequest struct {
	Id string `schema:"id" validate:"required,uuid"`
}

type deleteUserResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *deleteUserResponseData   `json:"data"`
}

type deleteUserResponseData struct {
}

func (a handler) checkPermission(ctx context.Context, permission string) *constants.ErrorResponse {
	claims, ok := ctx.Value(constants.ClaimsContextKey).(*objects.JWTClaims)
	if !ok || claims == nil {
		return constants.ErrTokenInvalid
	}
	fmt.Print("------------claims.Id")
	fmt.Print(claims.Id)
	fmt.Print("------------claims.Id")
	if !claims.IsSuperAdmin {
		if !utils.InArrayExist(permission, claims.Permissions) {
			return constants.ErrIneligibleAccess
		}
	}

	return nil
}
