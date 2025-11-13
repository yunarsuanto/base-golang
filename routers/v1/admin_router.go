package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yunarsuanto/base-go/infra/initiator/handler"
	"github.com/yunarsuanto/base-go/infra/middleware"
)

func adminRouter(handlerCtx *handler.HandlerCtx, mw middleware.MiddlewareInterface, r *mux.Router) {
	a := r.PathPrefix("/admin").Subrouter()
	a.Use(mw.GeneralAccessToken)

	a.HandleFunc("/user/get", handlerCtx.UserHandler.ListUser).Methods(http.MethodGet)
	a.HandleFunc("/user/detail/{id}", handlerCtx.UserHandler.DetailUser).Methods(http.MethodGet)
	a.HandleFunc("/user/create", handlerCtx.UserHandler.CreateUser).Methods(http.MethodPost)
	a.HandleFunc("/user/update", handlerCtx.UserHandler.UpdateUser).Methods(http.MethodPatch)
	a.HandleFunc("/user/delete", handlerCtx.UserHandler.DeleteUser).Methods(http.MethodDelete)

	a.HandleFunc("/role/get", handlerCtx.RoleHandler.ListRole).Methods(http.MethodGet)
	a.HandleFunc("/role/detail/{id}", handlerCtx.RoleHandler.DetailRole).Methods(http.MethodGet)
	a.HandleFunc("/role/create", handlerCtx.RoleHandler.CreateRole).Methods(http.MethodPost)
	a.HandleFunc("/role/update", handlerCtx.RoleHandler.UpdateRole).Methods(http.MethodPatch)
	a.HandleFunc("/role/delete", handlerCtx.RoleHandler.DeleteRole).Methods(http.MethodDelete)

	a.HandleFunc("/permission/get", handlerCtx.PermissionHandler.ListPermission).Methods(http.MethodGet)
	a.HandleFunc("/permission/create", handlerCtx.PermissionHandler.CreatePermission).Methods(http.MethodPost)
	a.HandleFunc("/permission/update", handlerCtx.PermissionHandler.UpdatePermission).Methods(http.MethodPatch)
	a.HandleFunc("/permission/delete", handlerCtx.PermissionHandler.DeletePermission).Methods(http.MethodDelete)

	a.HandleFunc("/role-permission/upsert", handlerCtx.RolePermissionHandler.UpsertRolePermission).Methods(http.MethodPatch)
	a.HandleFunc("/role-permission/delete", handlerCtx.RolePermissionHandler.DeleteRolePermission).Methods(http.MethodDelete)
	a.HandleFunc("/user-role/upsert", handlerCtx.UserRoleHandler.UpsertUserRole).Methods(http.MethodPatch)
	a.HandleFunc("/user-role/delete", handlerCtx.UserRoleHandler.DeleteUserRole).Methods(http.MethodDelete)

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

	a.HandleFunc("/lesson/copy", handlerCtx.LessonHandler.CopyLessonItem).Methods(http.MethodPost)

	a.HandleFunc("/lesson-item/get", handlerCtx.LessonItemHandler.ListLessonItem).Methods(http.MethodGet)
	a.HandleFunc("/lesson-item/detail/{id}", handlerCtx.LessonItemHandler.DetailLessonItem).Methods(http.MethodGet)
	a.HandleFunc("/lesson-item/create", handlerCtx.LessonItemHandler.CreateLessonItem).Methods(http.MethodPost)
	a.HandleFunc("/lesson-item/update", handlerCtx.LessonItemHandler.UpdateLessonItem).Methods(http.MethodPatch)
	a.HandleFunc("/lesson-item/delete", handlerCtx.LessonItemHandler.DeleteLessonItem).Methods(http.MethodDelete)

	// a.HandleFunc("/location/province/get", handlerCtx.GeneralLocationHandler.GetListProvince).Methods(http.MethodGet)
	// a.HandleFunc(utils.ParsePath("/notification/%s/detail", constants.IdPathVariable), handlerCtx.GeneralNotificationHandler.GetDetail).Methods(http.MethodGet)
}
