package user_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
)

type UserRepositoryInterface interface {
	GetList(ctx context.Context, tx *sqlx.Tx, pagination *objects.Pagination, req objects.ListUserRequest) ([]models.GetUser, *constants.ErrorResponse)
	GetById(ctx context.Context, tx *sqlx.Tx, userId string) (models.GetUser, *constants.ErrorResponse)
	GetByEmployeeId(ctx context.Context, tx *sqlx.Tx, employeeId string) (models.GetUser, *constants.ErrorResponse)
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateUser) *constants.ErrorResponse
	Update(ctx context.Context, tx *sqlx.Tx, data models.UpdateUser) *constants.ErrorResponse
	Delete(ctx context.Context, tx *sqlx.Tx, id string) *constants.ErrorResponse
	GetByEmail(ctx context.Context, tx *sqlx.Tx, email, username, excludingId string) (models.GetUser, *constants.ErrorResponse)
	GetByIdNumber(ctx context.Context, tx *sqlx.Tx, idNumber, excludingId string) (models.GetUser, *constants.ErrorResponse)
	GetByPhoneNumber(ctx context.Context, tx *sqlx.Tx, phoneNumber, excludingId string) (models.GetUser, *constants.ErrorResponse)
	ChangePassword(ctx context.Context, tx *sqlx.Tx, data models.ChangeUserPassword) *constants.ErrorResponse
	UpdateProfile(ctx context.Context, tx *sqlx.Tx, data models.UpdateUser) *constants.ErrorResponse
	UpdateActivation(ctx context.Context, tx *sqlx.Tx, id string, isActive bool) *constants.ErrorResponse
}

func NewUserRepository() UserRepositoryInterface {
	return &userRepository{}
}
