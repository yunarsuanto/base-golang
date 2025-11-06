package permission_service

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

func (a service) ListPermission(ctx context.Context, pagination *objects.Pagination) ([]objects.ListPermissionResponse, *constants.ErrorResponse) {
	var result []objects.ListPermissionResponse

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	resultData, errs := a.PermissionRepo.ListPermission(ctx, tx, pagination)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	for _, v := range resultData {
		result = append(result, objects.ListPermissionResponse(v))
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}

func (a service) CreatePermission(ctx context.Context, req objects.CreatePermissionRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	createData := models.CreatePermission{
		Name: req.Name,
	}

	errs := a.PermissionRepo.CreatePermission(ctx, tx, createData)
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
func (a service) UpdatePermission(ctx context.Context, req objects.UpdatePermissionRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	updateData := models.UpdatePermission{
		Id:   req.Id,
		Name: req.Name,
	}

	errs := a.PermissionRepo.UpdatePermission(ctx, tx, updateData)
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
func (a service) DeletePermission(ctx context.Context, req objects.DeletePermissionRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	deleteData := models.DeletePermission{
		Id: req.Id,
	}

	errs := a.PermissionRepo.DeletePermission(ctx, tx, deleteData)
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
