package user_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

type UserRepositoryInterface interface {
	GetByUsername(ctx context.Context, tx *sqlx.Tx, username string) (models.ListUser, *constants.ErrorResponse)
	GetById(ctx context.Context, tx *sqlx.Tx, id string) (models.ListUser, *constants.ErrorResponse)
	GetByTokenVerification(ctx context.Context, tx *sqlx.Tx, tokenVerification string) (models.ListUser, *constants.ErrorResponse)

	ListUser(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination) ([]models.ListUser, *constants.ErrorResponse)
	DetailUser(ctx context.Context, tx *sqlx.Tx, id string) ([]models.DetailUser, *constants.ErrorResponse)
	CreateUser(ctx context.Context, tx *sqlx.Tx, data models.CreateUser) *constants.ErrorResponse
	UpdateUser(ctx context.Context, tx *sqlx.Tx, data models.UpdateUser) *constants.ErrorResponse
	DeleteUser(ctx context.Context, tx *sqlx.Tx, data models.DeleteUser) *constants.ErrorResponse

	UpdateTokenVerification(ctx context.Context, tx *sqlx.Tx, data models.UpdateUserTokenVerification) *constants.ErrorResponse
	UpdateTokenVerificationIsActiveUser(ctx context.Context, tx *sqlx.Tx, data models.UpdateUserIsActiveTokenVerification) *constants.ErrorResponse

	CreateUserProfile(ctx context.Context, tx *sqlx.Tx, data models.CreateUserProfile) *constants.ErrorResponse
}

func NewUserRepository() UserRepositoryInterface {
	return &repository{}
}
