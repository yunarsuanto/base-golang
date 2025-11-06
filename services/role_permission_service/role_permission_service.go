package role_permission_service

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

func (a service) UpsertRolePermission(ctx context.Context, req objects.UpsertRolePermissionRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	createData := models.UpsertRolePermissionRequest{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	}

	errs := a.RolePermissionRepo.UpsertRolePermission(ctx, tx, createData)
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

func (a service) DeleteRolePermission(ctx context.Context, req objects.DeleteRolePermissionRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	createData := models.DeleteRolePermissionRequest{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	}

	errs := a.RolePermissionRepo.DeleteRolePermission(ctx, tx, createData)
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
