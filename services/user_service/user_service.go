package user_service

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type service struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (a service) ListUser(ctx context.Context, pagination *objects.Pagination) ([]objects.ListUserResponse, *constants.ErrorResponse) {
	var result []objects.ListUserResponse

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	resultData, errs := a.UserRepo.ListUser(ctx, tx, pagination)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	for _, v := range resultData {
		result = append(result, objects.ListUserResponse(v))
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a service) DetailUser(ctx context.Context, req objects.DetailUserRequest) ([]objects.DetailUserResponse, *constants.ErrorResponse) {
	var result []objects.DetailUserResponse

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	resultData, errs := a.UserRepo.DetailUser(ctx, tx, req.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	for _, v := range resultData {
		result = append(result, objects.DetailUserResponse(v))
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a service) CreateUser(ctx context.Context, req objects.CreateUserRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	pass, err := utils.HashPassword(req.Password)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}
	createData := models.CreateUser{
		Username: req.Username,
		Password: pass,
		IsActive: true,
	}

	errs := a.UserRepo.CreateUser(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return utils.ErrorInternalServer(err.Error())
	}

	return nil
}
func (a service) UpdateUser(ctx context.Context, req objects.UpdateUserRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	updateData := models.UpdateUser{
		Id:       req.Id,
		Username: req.Username,
	}

	errs := a.UserRepo.UpdateUser(ctx, tx, updateData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return utils.ErrorInternalServer(err.Error())
	}

	return nil
}
func (a service) DeleteUser(ctx context.Context, req objects.DeleteUserRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	deleteData := models.DeleteUser{
		Id: req.Id,
	}

	errs := a.UserRepo.DeleteUser(ctx, tx, deleteData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return utils.ErrorInternalServer(err.Error())
	}

	return nil
}
