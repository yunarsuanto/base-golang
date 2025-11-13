package category_lesson_handler

import (
	"github.com/yunarsuanto/base-go/constants"
	common_input_handler "github.com/yunarsuanto/base-go/handler"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
	"golang.org/x/net/context"
)

type listCategoryLessonRequest struct {
	common_input_handler.PaginationRequest
}

type listCategoryLessonResponse struct {
	Meta       common_input_handler.Meta         `json:"meta"`
	Pagination *common_input_handler.Pagination  `json:"pagination"`
	Data       []*listCategoryLessonResponseData `json:"data"`
}

type listCategoryLessonResponseData struct {
	Id                 string `json:"id"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	CategoryLessonType string `json:"category_lesson_type"`
	Media              string `json:"media"`
}

type detailCategoryLessonRequest struct {
	Id string `json:"id"`
}

type detailCategoryLessonResponse struct {
	Meta common_input_handler.Meta        `json:"meta"`
	Data detailCategoryLessonDataResponse `json:"data"`
}

type detailCategoryLessonDataResponse struct {
	Id                 string `json:"id"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	CategoryLessonType string `json:"category_lesson_type"`
	Media              string `json:"media"`
}

type createCategoryLessonRequest struct {
	Title              string `json:"title" schema:"title" validate:"required"`
	Description        string `json:"description" schema:"description" validate:"required"`
	CategoryLessonType string `json:"category_lesson_type" schema:"category_lesson_type" validate:"required"`
	Media              string `json:"media" schema:"media" validate:"required"`
}

type createCategoryLessonResponse struct {
	Meta common_input_handler.Meta         `json:"meta"`
	Data *createCategoryLessonResponseData `json:"data"`
}

type createCategoryLessonResponseData struct {
}

type updateCategoryLessonRequest struct {
	Id                 string `json:"id" schema:"id" validate:"required,uuid"`
	Title              string `json:"title" schema:"title" validate:"required"`
	Description        string `json:"description" schema:"description" validate:"required"`
	CategoryLessonType string `json:"category_lesson_type" schema:"category_lesson_type" validate:"required"`
	Media              string `json:"media" schema:"media" validate:"required"`
}

type updateCategoryLessonResponse struct {
	Meta common_input_handler.Meta         `json:"meta"`
	Data *updateCategoryLessonResponseData `json:"data"`
}

type updateCategoryLessonResponseData struct {
}

type deleteCategoryLessonRequest struct {
	Id string `json:"id" schema:"id" validate:"required,uuid"`
}

type deleteCategoryLessonResponse struct {
	Meta common_input_handler.Meta         `json:"meta"`
	Data *deleteCategoryLessonResponseData `json:"data"`
}

type deleteCategoryLessonResponseData struct {
}

type listCategoryLessonPublicResponse struct {
	Meta common_input_handler.Meta            `json:"meta"`
	Data listCategoryLessonPublicResponseData `json:"data"`
}

type listCategoryLessonPublicResponseData struct {
	Id                 string                                       `json:"id"`
	Title              string                                       `json:"title"`
	Description        string                                       `json:"description"`
	CategoryLessonType string                                       `json:"category_lesson_type"`
	Media              string                                       `json:"media"`
	Lessons            []listCategoryLessonPublicResponseDataLesson `json:"lessons"`
}

type listCategoryLessonPublicResponseDataLesson struct {
	Id                 string                                           `json:"id"`
	Title              string                                           `json:"title"`
	Description        string                                           `json:"description"`
	CategoryLessonType string                                           `json:"category_lesson_type"`
	Media              string                                           `json:"media"`
	Level              uint32                                           `json:"level"`
	Items              []listCategoryLessonPublicResponseDataLessonItem `json:"items"`
}

type listCategoryLessonPublicResponseDataLessonItem struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Order   uint32 `json:"order"`
	Media   string `json:"media"`
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
