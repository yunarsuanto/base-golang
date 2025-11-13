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

	a.HandleFunc("/lesson/get", handlerCtx.LessonHandler.ListLesson).Methods(http.MethodGet)
	a.HandleFunc("/lesson/detail/{id}", handlerCtx.LessonHandler.DetailLesson).Methods(http.MethodGet)

	a.HandleFunc("/lesson-item/get", handlerCtx.LessonItemHandler.ListLessonItem).Methods(http.MethodGet)
	a.HandleFunc("/lesson-item/detail/{id}", handlerCtx.LessonItemHandler.DetailLessonItem).Methods(http.MethodGet)

	// a.HandleFunc("/location/province/get", handlerCtx.GeneralLocationHandler.GetListProvince).Methods(http.MethodGet)
	// a.HandleFunc(utils.ParsePath("/notification/%s/detail", constants.IdPathVariable), handlerCtx.GeneralNotificationHandler.GetDetail).Methods(http.MethodGet)
}
