package utils

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/objects"
)

// GetJwtClaimsFromContext function to get jwt claims from context
func GetJwtClaimsFromContext(ctx context.Context) (*objects.JWTClaims, *constants.ErrorResponse) {
	claims, ok := ctx.Value(constants.ClaimsContextKey).(*objects.JWTClaims)
	if !ok {
		return nil, constants.ErrInvalidClaims
	}
	return claims, nil
}
