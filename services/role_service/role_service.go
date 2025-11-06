package role_service

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

func (a service) ListRole(ctx context.Context, pagination *objects.Pagination) ([]objects.ListRoleResponse, *constants.ErrorResponse) {
	var result []objects.ListRoleResponse

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	resultData, errs := a.RoleRepo.ListRole(ctx, tx, pagination)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	for _, v := range resultData {
		result = append(result, objects.ListRoleResponse(v))
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a service) DetailRole(ctx context.Context, req objects.DetailRoleRequest) ([]objects.DetailRoleResponse, *constants.ErrorResponse) {
	var result []objects.DetailRoleResponse

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	resultData, errs := a.RoleRepo.DetailRole(ctx, tx, req.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	for _, v := range resultData {
		result = append(result, objects.DetailRoleResponse(v))
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a service) CreateRole(ctx context.Context, req objects.CreateRoleRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	createData := models.CreateRole{
		Name: req.Name,
	}

	errs := a.RoleRepo.CreateRole(ctx, tx, createData)
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
func (a service) UpdateRole(ctx context.Context, req objects.UpdateRoleRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	updateData := models.UpdateRole{
		Id:   req.Id,
		Name: req.Name,
	}

	errs := a.RoleRepo.UpdateRole(ctx, tx, updateData)
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
func (a service) DeleteRole(ctx context.Context, req objects.DeleteRoleRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	deleteData := models.DeleteRole{
		Id: req.Id,
	}

	errs := a.RoleRepo.DeleteRole(ctx, tx, deleteData)
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
