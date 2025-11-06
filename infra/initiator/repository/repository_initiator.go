package repository

import (
	"github.com/yunarsuanto/base-go/repositories/activity_log_repository"
	"github.com/yunarsuanto/base-go/repositories/category_lesson_repository"
	"github.com/yunarsuanto/base-go/repositories/lesson_item_repository"
	"github.com/yunarsuanto/base-go/repositories/lesson_repository"
	"github.com/yunarsuanto/base-go/repositories/permission_repository"
	"github.com/yunarsuanto/base-go/repositories/role_permission_repository"
	"github.com/yunarsuanto/base-go/repositories/role_repository"
	"github.com/yunarsuanto/base-go/repositories/user_repository"
	"github.com/yunarsuanto/base-go/repositories/user_role_repository"
)

type RepoCtx struct {
	UserRepo           user_repository.UserRepositoryInterface
	UserRoleRepo       user_role_repository.UsersRolesRepositoryInterface
	RolePermissionRepo role_permission_repository.RolesPermissionsRepositoryInterface
	ActivityLogRepo    activity_log_repository.ActivityLogRepositoryInterface
	RoleRepo           role_repository.RoleRepositoryInterface
	PermissionRepo     permission_repository.PermissionRepositoryInterface
	CategoryLessonRepo category_lesson_repository.CategoryLessonRepositoryInterface
	LessonRepo         lesson_repository.LessonRepositoryInterface
	LessonItemRepo     lesson_item_repository.LessonItemRepositoryInterface
}

func InitRepoCtx() *RepoCtx {
	return &RepoCtx{
		UserRepo:           user_repository.NewUserRepository(),
		UserRoleRepo:       user_role_repository.NewUsersRolesRepository(),
		RolePermissionRepo: role_permission_repository.NewRolesPermissionsRepository(),
		ActivityLogRepo:    activity_log_repository.NewActivityLogRepository(),
		RoleRepo:           role_repository.NewRoleRepository(),
		PermissionRepo:     permission_repository.NewPermissionRepository(),
		CategoryLessonRepo: category_lesson_repository.NewCategoryLessonRepository(),
		LessonRepo:         lesson_repository.NewLessonRepository(),
		LessonItemRepo:     lesson_item_repository.NewLessonItemRepository(),
	}
}
