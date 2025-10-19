package utils

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
)

func VerifyPermission(ctx context.Context, permission string) *constants.ErrorResponse {
	claims, errs := GetJwtClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	hasPermission := InArrayExist(permission, claims.Permissions)
	if !hasPermission {
		return constants.ErrIneligibleAccess
	}

	return nil
}
