package user_role_service

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

func (a service) UpdateUserRole(ctx context.Context, req objects.UpsertUserRoleRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	updateData := models.UpsertUserRoleRequest{
		UserId:   req.UserId,
		RoleId:   req.RoleId,
		IsActive: req.IsActive,
	}

	errs := a.UserRoleRepo.UpsertUserRole(ctx, tx, updateData)
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

func (a service) DeleteUserRole(ctx context.Context, req objects.DeleteUserRoleRequest) *constants.ErrorResponse {
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return utils.ErrorInternalServer(err.Error())
	}

	updateData := models.DeleteUserRoleRequest{
		UserId: req.UserId,
		RoleId: req.RoleId,
	}

	errs := a.UserRoleRepo.DeleteUserRole(ctx, tx, updateData)
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
