package handler

import (
	"github.com/yunarsuanto/base-go/handler/auth_handler"
	"github.com/yunarsuanto/base-go/handler/category_lesson_handler"
	"github.com/yunarsuanto/base-go/handler/file_handler"
	"github.com/yunarsuanto/base-go/handler/lesson_handler"
	"github.com/yunarsuanto/base-go/handler/lesson_item_handler"
	"github.com/yunarsuanto/base-go/handler/permission_handler"
	"github.com/yunarsuanto/base-go/handler/role_handler"
	"github.com/yunarsuanto/base-go/handler/role_permission_handler"
	"github.com/yunarsuanto/base-go/handler/user_handler"
	"github.com/yunarsuanto/base-go/handler/user_role_handler"
	"github.com/yunarsuanto/base-go/infra/initiator/service"
)

type HandlerCtx struct {
	FileHandler           file_handler.FileHandlerInterface
	AuthHandler           auth_handler.AuthHandlerInterface
	UserHandler           user_handler.UserHandlerInterface
	RoleHandler           role_handler.RoleHandlerInterface
	PermissionHandler     permission_handler.PermissionHandlerInterface
	RolePermissionHandler role_permission_handler.RolePermissionHandlerInterface
	UserRoleHandler       user_role_handler.UserRoleHandlerInterface
	CategoryLessonHandler category_lesson_handler.CategoryLessonHandlerInterface
	LessonHandler         lesson_handler.LessonHandlerInterface
	LessonItemHandler     lesson_item_handler.LessonItemHandlerInterface
}

func InitHandlerCtx(serviceCtx *service.ServiceCtx) *HandlerCtx {
	return &HandlerCtx{
		FileHandler:           file_handler.NewFileHandler(serviceCtx),
		AuthHandler:           auth_handler.NewAuthHandler(serviceCtx),
		UserHandler:           user_handler.NewUserHandler(serviceCtx),
		RoleHandler:           role_handler.NewRoleHandler(serviceCtx),
		PermissionHandler:     permission_handler.NewPermissionHandler(serviceCtx),
		RolePermissionHandler: role_permission_handler.NewRolePermissionHandler(serviceCtx),
		UserRoleHandler:       user_role_handler.NewUserRoleHandler(serviceCtx),
		CategoryLessonHandler: category_lesson_handler.NewCategoryLessonHandler(serviceCtx),
		LessonHandler:         lesson_handler.NewLessonHandler(serviceCtx),
		LessonItemHandler:     lesson_item_handler.NewLessonItemHandler(serviceCtx),
	}
}
