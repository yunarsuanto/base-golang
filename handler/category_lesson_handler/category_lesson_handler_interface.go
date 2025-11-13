package category_lesson_handler

import (
	"net/http"

	"github.com/yunarsuanto/base-go/infra/initiator/service"
)

type CategoryLessonHandlerInterface interface {
	ListCategoryLessonPublic(w http.ResponseWriter, r *http.Request)

	ListCategoryLesson(w http.ResponseWriter, r *http.Request)
	DetailCategoryLesson(w http.ResponseWriter, r *http.Request)
	CreateCategoryLesson(w http.ResponseWriter, r *http.Request)
	UpdateCategoryLesson(w http.ResponseWriter, r *http.Request)
	DeleteCategoryLesson(w http.ResponseWriter, r *http.Request)
}

func NewCategoryLessonHandler(serviceCtx *service.ServiceCtx) CategoryLessonHandlerInterface {
	return &handler{
		serviceCtx,
	}
}
