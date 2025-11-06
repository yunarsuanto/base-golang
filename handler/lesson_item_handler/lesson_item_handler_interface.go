package lesson_item_handler

import (
	"net/http"

	"github.com/yunarsuanto/base-go/infra/initiator/service"
)

type LessonItemHandlerInterface interface {
	ListLessonItem(w http.ResponseWriter, r *http.Request)
	DetailLessonItem(w http.ResponseWriter, r *http.Request)
	CreateLessonItem(w http.ResponseWriter, r *http.Request)
	UpdateLessonItem(w http.ResponseWriter, r *http.Request)
	DeleteLessonItem(w http.ResponseWriter, r *http.Request)
}

func NewLessonItemHandler(serviceCtx *service.ServiceCtx) LessonItemHandlerInterface {
	return &handler{
		serviceCtx,
	}
}
