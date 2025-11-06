package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yunarsuanto/base-go/infra/initiator/handler"
	"github.com/yunarsuanto/base-go/infra/middleware"
)

func userRouter(handlerCtx *handler.HandlerCtx, mw middleware.MiddlewareInterface, r *mux.Router) {
	a := r.PathPrefix("/user").Subrouter()
	a.Use(mw.GeneralAccessToken)

	a.HandleFunc("/category-lesson/get", handlerCtx.CategoryLessonHandler.ListCategoryLesson).Methods(http.MethodGet)
	a.HandleFunc("/category-lesson/detail/{id}", handlerCtx.CategoryLessonHandler.DetailCategoryLesson).Methods(http.MethodGet)
	a.HandleFunc("/category-lesson/create", handlerCtx.CategoryLessonHandler.CreateCategoryLesson).Methods(http.MethodPost)
	a.HandleFunc("/category-lesson/update", handlerCtx.CategoryLessonHandler.UpdateCategoryLesson).Methods(http.MethodPatch)
	a.HandleFunc("/category-lesson/delete", handlerCtx.CategoryLessonHandler.DeleteCategoryLesson).Methods(http.MethodDelete)

	a.HandleFunc("/lesson/get", handlerCtx.LessonHandler.ListLesson).Methods(http.MethodGet)
	a.HandleFunc("/lesson/detail/{id}", handlerCtx.LessonHandler.DetailLesson).Methods(http.MethodGet)
	a.HandleFunc("/lesson/create", handlerCtx.LessonHandler.CreateLesson).Methods(http.MethodPost)
	a.HandleFunc("/lesson/update", handlerCtx.LessonHandler.UpdateLesson).Methods(http.MethodPatch)
	a.HandleFunc("/lesson/delete", handlerCtx.LessonHandler.DeleteLesson).Methods(http.MethodDelete)

	a.HandleFunc("/lesson-item/get", handlerCtx.LessonItemHandler.ListLessonItem).Methods(http.MethodGet)
	a.HandleFunc("/lesson-item/detail/{id}", handlerCtx.LessonItemHandler.DetailLessonItem).Methods(http.MethodGet)
	a.HandleFunc("/lesson-item/create", handlerCtx.LessonItemHandler.CreateLessonItem).Methods(http.MethodPost)
	a.HandleFunc("/lesson-item/update", handlerCtx.LessonItemHandler.UpdateLessonItem).Methods(http.MethodPatch)
	a.HandleFunc("/lesson-item/delete", handlerCtx.LessonItemHandler.DeleteLessonItem).Methods(http.MethodDelete)

	// a.HandleFunc("/location/province/get", handlerCtx.GeneralLocationHandler.GetListProvince).Methods(http.MethodGet)
	// a.HandleFunc(utils.ParsePath("/notification/%s/detail", constants.IdPathVariable), handlerCtx.GeneralNotificationHandler.GetDetail).Methods(http.MethodGet)
}
