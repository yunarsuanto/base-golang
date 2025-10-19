package auth_service

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

func (a service) getUserPermissions(ctx context.Context, tx *sqlx.Tx, platform, userId string) ([]models.GetPermissionCode, *constants.ErrorResponse) {
	var result []models.GetPermissionCode

	var permissionCode string
	switch platform {
	case constants.WebPlatform:
		permissionCode = constants.PermissionAccessWeb
	case constants.MobilePlatform:
		permissionCode = constants.PermissionAccessMobile
	default:
		_ = tx.Rollback()
		return result, constants.ErrIneligibleAccess
	}
	roleData, errs := a.UserRoleRepo.GetList(ctx, tx, objects.NewPagination().AllData(), objects.ListUsersRolesRequest{UserIds: []string{userId}})
	if errs != nil {
		return result, errs
	}
	if len(roleData) == 0 {
		return result, constants.ErrIneligibleAccess
	}

	roleIds := make([]string, len(roleData))
	for i, v := range roleData {
		roleIds[i] = v.RoleId
	}

	_, errs = a.RolePermissionRepo.GetByRoleIdsPermissionCode(ctx, tx, roleIds, permissionCode)
	if errs != nil {
		return result, errs
	}

	result, errs = a.RolePermissionRepo.GetDistinctPermissionByRoleIds(ctx, tx, roleIds)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	return result, nil
}

func (a service) generateToken(ctx context.Context, tx *sqlx.Tx, platform, fcmToken string, userData models.GetUser, permissionData []models.GetPermissionCode) (objects.Login, *constants.ErrorResponse) {
	var result objects.Login

	permissions := make([]string, len(permissionData))
	for i, v := range permissionData {
		permissions[i] = v.PermissionCode
	}

	accessTokenData := objects.JWTRequest{
		Id:          userData.Id,
		Platform:    platform,
		Permissions: permissions,
	}
	accessTokenDuration := time.Duration(a.Config.JwtConfig.AccessTokenDuration) * time.Minute
	refreshTokenData := objects.JWTRequest{
		Id:       userData.Id,
		Platform: platform,
	}
	refreshTokenDuration := time.Duration(a.Config.JwtConfig.RefreshTokenDuration) * time.Hour * 24 * 30

	accessToken, errs := a.Jwt.GenerateJWTToken(ctx, accessTokenData, accessTokenDuration, constants.AppName, constants.RedisKeyAccessToken)
	if errs != nil {
		return result, errs
	}
	refreshToken, errs := a.Jwt.GenerateJWTToken(ctx, refreshTokenData, refreshTokenDuration, constants.AppName, constants.RedisKeyRefreshToken)
	if errs != nil {
		return result, errs
	}

	errs = a.UserTokenRepo.Upsert(ctx, tx, models.UpsertUserToken{
		UserId:     userData.Id,
		Platform:   platform,
		FcmToken:   fcmToken,
		ExpiryTime: time.Now().Add(refreshTokenDuration),
	})
	if errs != nil {
		return result, errs
	}

	result = objects.Login{
		AccessToken:   accessToken,
		RefreshToken:  refreshToken,
		ExpiredAt:     time.Now().Add(accessTokenDuration),
		IsReporter:    userData.IsReporter,
		IsVerificator: userData.IsVerificator,
		Permissions:   permissions,
	}

	return result, nil
}
