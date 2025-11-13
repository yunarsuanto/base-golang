package lesson_item_handler

import (
	"github.com/yunarsuanto/base-go/constants"
	common_input_handler "github.com/yunarsuanto/base-go/handler"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
	"golang.org/x/net/context"
)

type listLessonItemRequest struct {
	common_input_handler.PaginationRequest
	LessonId string `json:"lesson_id" schema:"lesson_id" validate:"omitempty,required,uuid"`
}

type listLessonItemResponse struct {
	Meta       common_input_handler.Meta        `json:"meta"`
	Pagination *common_input_handler.Pagination `json:"pagination"`
	Data       []*listLessonItemResponseData    `json:"data"`
}

type listLessonItemResponseData struct {
	Id       string `json:"id"`
	LessonId string `json:"lesson_id"`
	Content  string `json:"content"`
	Order    uint32 `json:"order"`
	Media    string `json:"media"`
	Group    uint32 `json:"group"`
	IsDone   bool   `json:"is_done"`
}

type detailLessonItemRequest struct {
	Id string `json:"id" schema:"id" validate:"omitempty,required,uuid"`
}

type detailLessonItemResponse struct {
	Meta common_input_handler.Meta    `json:"meta"`
	Data detailLessonItemResponseData `json:"data"`
}

type detailLessonItemResponseData struct {
	Id       string `json:"id"`
	LessonId string `json:"lesson_id"`
	Content  string `json:"content"`
	Order    uint32 `json:"order"`
	Media    string `json:"media"`
	Group    uint32 `json:"group"`
	IsDone   bool   `json:"is_done"`
}

type createLessonItemRequest struct {
	LessonId string `json:"lesson_id" schema:"lesson_id" validate:"required"`
	Content  string `json:"content" schema:"content" validate:"required"`
	Order    uint32 `json:"order" schema:"order" validate:"required"`
	Media    string `json:"media" schema:"media" validate:"required"`
	Group    uint32 `json:"group" schema:"group" validate:"required"`
	IsDone   bool   `json:"is_done" schema:"is_done"`
}

type createLessonItemResponse struct {
	Meta common_input_handler.Meta     `json:"meta"`
	Data *createLessonItemResponseData `json:"data"`
}

type createLessonItemResponseData struct {
}

type updateLessonItemRequest struct {
	Id       string `json:"id" schema:"id" validate:"required,uuid"`
	LessonId string `json:"lesson_id" schema:"lesson_id" validate:"required"`
	Content  string `json:"content" schema:"content" validate:"required"`
	Order    uint32 `json:"order" schema:"order" validate:"required"`
	Media    string `json:"media" schema:"media" validate:"required"`
	Group    uint32 `json:"group" schema:"group" validate:"required"`
	IsDone   bool   `json:"is_done" schema:"is_done"`
}

type updateLessonItemResponse struct {
	Meta common_input_handler.Meta     `json:"meta"`
	Data *updateLessonItemResponseData `json:"data"`
}

type updateLessonItemResponseData struct {
}

type deleteLessonItemRequest struct {
	Id string `json:"id" schema:"id" validate:"required,uuid"`
}

type deleteLessonItemResponse struct {
	Meta common_input_handler.Meta     `json:"meta"`
	Data *deleteLessonItemResponseData `json:"data"`
}

type deleteLessonItemResponseData struct {
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
