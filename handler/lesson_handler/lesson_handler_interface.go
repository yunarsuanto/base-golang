package lesson_handler

import (
	"net/http"

	"github.com/yunarsuanto/base-go/infra/initiator/service"
)

type LessonHandlerInterface interface {
	ListLesson(w http.ResponseWriter, r *http.Request)
	DetailLesson(w http.ResponseWriter, r *http.Request)
	CreateLesson(w http.ResponseWriter, r *http.Request)
	UpdateLesson(w http.ResponseWriter, r *http.Request)
	DeleteLesson(w http.ResponseWriter, r *http.Request)
	CopyLessonItem(w http.ResponseWriter, r *http.Request)
}

func NewLessonHandler(serviceCtx *service.ServiceCtx) LessonHandlerInterface {
	return &handler{
		serviceCtx,
	}
}
