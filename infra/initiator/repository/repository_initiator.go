package repository

import (
	"github.com/yunarsuanto/base-go/repositories/activity_log_repository"
	"github.com/yunarsuanto/base-go/repositories/role_permission_repository"
	"github.com/yunarsuanto/base-go/repositories/user_repository"
	"github.com/yunarsuanto/base-go/repositories/user_role_repository"
	"github.com/yunarsuanto/base-go/repositories/user_token_repository"
)

type RepoCtx struct {
	UserRepo           user_repository.UserRepositoryInterface
	UserRoleRepo       user_role_repository.UsersRolesRepositoryInterface
	RolePermissionRepo role_permission_repository.RolesPermissionsRepositoryInterface
	UserTokenRepo      user_token_repository.UserTokenRepositoryInterface
	ActivityLogRepo    activity_log_repository.ActivityLogRepositoryInterface
}

func InitRepoCtx() *RepoCtx {
	return &RepoCtx{
		UserRepo:           user_repository.NewUserRepository(),
		UserRoleRepo:       user_role_repository.NewUsersRolesRepository(),
		RolePermissionRepo: role_permission_repository.NewRolesPermissionsRepository(),
		UserTokenRepo:      user_token_repository.NewUserTokenRepository(),
		ActivityLogRepo:    activity_log_repository.NewActivityLogRepository(),
	}
}
