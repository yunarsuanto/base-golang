package lesson_handler

import (
	"github.com/yunarsuanto/base-go/constants"
	common_input_handler "github.com/yunarsuanto/base-go/handler"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
	"golang.org/x/net/context"
)

type listLessonRequest struct {
	common_input_handler.PaginationRequest
	CategoryLessonId string `schema:"category_lesson_id" validate:"omitempty,uuid"`
}

type listLessonResponse struct {
	Meta       common_input_handler.Meta        `json:"meta"`
	Pagination *common_input_handler.Pagination `json:"pagination"`
	Data       []*listLessonResponseData        `json:"data"`
}

type listLessonResponseData struct {
	Id                  string `json:"id"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	CategoryLessonId    string `json:"category_lesson_id"`
	LessonType          string `json:"lesson_type"`
	Media               string `json:"media"`
	CategoryLessonTitle string `json:"category_lesson_title"`
	Level               uint32 `json:"level"`
}

type detailLessonRequest struct {
	Id string `json:"id"`
}

type detailLessonResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data detailLessonResponseData  `json:"data"`
}

type detailLessonResponseData struct {
	Id               string `json:"id"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	CategoryLessonId string `json:"category_lesson_id"`
	LessonType       string `json:"lesson_type"`
	Media            string `json:"media"`
	Level            uint32 `json:"level"`
}

type createLessonRequest struct {
	Title            string `json:"title" schema:"title" validate:"required"`
	Description      string `json:"description" schema:"description" validate:"required"`
	CategoryLessonId string `json:"category_lesson_id" schema:"category_lesson_id" validate:"required"`
	LessonType       string `json:"lesson_type" schema:"lesson_type" validate:"required"`
	Media            string `json:"media" schema:"media" validate:"required"`
	Level            uint32 `json:"level" schema:"level" validate:"required"`
}

type createLessonResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *createLessonResponseData `json:"data"`
}

type createLessonResponseData struct {
}

type updateLessonRequest struct {
	Id               string `json:"id" schema:"id" validate:"required,uuid"`
	Title            string `json:"title" schema:"title" validate:"required"`
	Description      string `json:"description" schema:"description" validate:"required"`
	CategoryLessonId string `json:"category_lesson_id" schema:"category_lesson_id" validate:"required"`
	LessonType       string `json:"lesson_type" schema:"lesson_type" validate:"required"`
	Media            string `json:"media" schema:"media" validate:"required"`
	Level            uint32 `json:"level" schema:"level" validate:"required"`
}

type updateLessonResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *updateLessonResponseData `json:"data"`
}

type updateLessonResponseData struct {
}

type deleteLessonRequest struct {
	Id string `json:"id" schema:"id" validate:"required,uuid"`
}

type deleteLessonResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *deleteLessonResponseData `json:"data"`
}

type deleteLessonResponseData struct {
}

type copyLessonRequest struct {
	LessonId  string `json:"lesson_id" schema:"lesson_id" validate:"required,uuid"`
	Level     uint32 `json:"level" schema:"level" validate:"required"`
	LevelFrom uint32 `json:"level_from" schema:"level_from" validate:"required"`
}

type copyLessonResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *copyLessonResponseData   `json:"data"`
}

type copyLessonResponseData struct {
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
