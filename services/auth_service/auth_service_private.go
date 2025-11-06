package auth_service

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

func (a service) getAdminPermissions(ctx context.Context, tx *sqlx.Tx, userId string, isSuperAdmin bool) (string, []models.GetPermissionName, *constants.ErrorResponse) {
	var result []models.GetPermissionName

	// var permissionCode string
	// switch platform {
	// case constants.WebPlatform:
	// 	permissionCode = constants.PermissionAccessWeb
	// case constants.MobilePlatform:
	// 	permissionCode = constants.PermissionAccessMobile
	// default:
	// 	_ = tx.Rollback()
	// 	return result, constants.ErrIneligibleAccess
	// }

	roleData, errs := a.UserRoleRepo.GetByUserId(ctx, tx, objects.ListUserRoleRequest{UserId: userId})
	if errs != nil {
		return "", result, errs
	}

	// if roleData.Id == "" && !isSuperAdmin {
	// 	return result, constants.ErrIneligibleAccess
	// }

	if !isSuperAdmin && roleData.Id != "" {
		result, errs = a.RolePermissionRepo.GetDistinctPermissionByRoleId(ctx, tx, roleData.RoleId)
		if errs != nil {
			_ = tx.Rollback()
			return "", result, errs
		}
	}

	return roleData.RoleName, result, nil
}

func (a service) generateToken(ctx context.Context, platform, fcmToken string, userData models.ListUser, permissionData []models.GetPermissionName, isSuperAdmin bool, roleName string) (objects.LoginResponse, *constants.ErrorResponse) {
	var result objects.LoginResponse

	permissions := make([]string, len(permissionData))
	for i, v := range permissionData {
		permissions[i] = v.PermissionName
	}

	accessTokenData := objects.JWTRequest{
		Id:          userData.Id,
		Platform:    platform,
		Permissions: permissions,
		Role:        roleName,
	}
	accessTokenDuration := time.Duration(a.Config.JwtConfig.AccessTokenDuration) * time.Minute
	refreshTokenData := objects.JWTRequest{
		Id:       userData.Id,
		Platform: platform,
	}
	refreshTokenDuration := time.Duration(a.Config.JwtConfig.RefreshTokenDuration) * time.Hour * 24 * 30

	accessToken, errs := a.Jwt.GenerateJWTToken(ctx, accessTokenData, accessTokenDuration, constants.AppName, constants.RedisKeyAccessToken, isSuperAdmin)
	if errs != nil {
		return result, errs
	}
	refreshToken, errs := a.Jwt.GenerateJWTToken(ctx, refreshTokenData, refreshTokenDuration, constants.AppName, constants.RedisKeyRefreshToken, isSuperAdmin)
	if errs != nil {
		return result, errs
	}

	result = objects.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(accessTokenDuration),
		Permissions:  permissions,
		Role:         roleName,
	}

	return result, nil
}
