package service

import (
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/services/activity_log_service"
	"github.com/yunarsuanto/base-go/services/auth_service"
	"github.com/yunarsuanto/base-go/services/category_lesson_service"
	"github.com/yunarsuanto/base-go/services/file_service"
	"github.com/yunarsuanto/base-go/services/lesson_item_service"
	"github.com/yunarsuanto/base-go/services/lesson_service"
	"github.com/yunarsuanto/base-go/services/permission_service"
	"github.com/yunarsuanto/base-go/services/role_permission_service"
	"github.com/yunarsuanto/base-go/services/role_service"
	"github.com/yunarsuanto/base-go/services/user_role_service"
	"github.com/yunarsuanto/base-go/services/user_service"
)

type ServiceCtx struct {
	FileService           file_service.FileServiceInterface
	AuthService           auth_service.AuthServiceInterface
	ActivityLogService    activity_log_service.ActivityLogServiceInterface
	UserService           user_service.UserServiceInterface
	RoleService           role_service.RoleServiceInterface
	PermissionService     permission_service.PermissionServiceInterface
	RolePermissionService role_permission_service.RolePermissionServiceInterface
	UserRoleService       user_role_service.UserRoleServiceInterface
	CategoryLessonService category_lesson_service.CategoryLessonServiceInterface
	LessonService         lesson_service.LessonServiceInterface
	LessonItemService     lesson_item_service.LessonItemServiceInterface
}

func InitServiceCtx(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) *ServiceCtx {
	return &ServiceCtx{
		FileService:           file_service.NewFileService(repoCtx, infraCtx),
		AuthService:           auth_service.NewAuthService(repoCtx, infraCtx),
		ActivityLogService:    activity_log_service.NewActivityLogService(repoCtx, infraCtx),
		UserService:           user_service.NewUserService(repoCtx, infraCtx),
		RoleService:           role_service.NewRoleService(repoCtx, infraCtx),
		PermissionService:     permission_service.NewPermissionService(repoCtx, infraCtx),
		RolePermissionService: role_permission_service.NewRolePermissionService(repoCtx, infraCtx),
		UserRoleService:       user_role_service.NewUserRoleService(repoCtx, infraCtx),
		CategoryLessonService: category_lesson_service.NewCategoryLessonService(repoCtx, infraCtx),
		LessonService:         lesson_service.NewLessonService(repoCtx, infraCtx),
		LessonItemService:     lesson_item_service.NewLessonItemService(repoCtx, infraCtx),
	}
}
