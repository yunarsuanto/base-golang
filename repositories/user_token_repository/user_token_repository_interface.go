package user_token_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

type UserTokenRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination, req objects.ListUserTokenRequest) ([]models.GetUserToken, *constants.ErrorResponse)
	Upsert(ctx context.Context, tx *sqlx.Tx, data models.UpsertUserToken) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, userId, platform string) *constants.ErrorResponse
}

func NewUserTokenRepository() UserTokenRepositoryInterface {
	return &userTokenRepository{}
}
