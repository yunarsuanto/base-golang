package auth_service

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type service struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (a service) Login(ctx context.Context, data objects.LoginRequest) (objects.Login, *constants.ErrorResponse) {
	var result objects.Login

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	userData, errs := a.UserRepo.GetByEmail(ctx, tx, data.Email, "", "")
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	if !userData.IsActive {
		_ = tx.Rollback()
		return result, constants.ErrInactiveUser
	}

	// hospitalData, errs := a.UsersHospitalsRepo.GetByUserIds(ctx, tx, []string{userData.Id})
	// if errs != nil {
	// 	_ = tx.Rollback()
	// 	return result, errs
	// }

	permissionData, errs := a.getUserPermissions(ctx, tx, data.Platform, userData.Id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	isCheckPassword := utils.CheckPasswordHash(data.Password, userData.Password)
	if !isCheckPassword {
		_ = tx.Rollback()
		return result, constants.ErrEmailAndPasswordNotMatch
	}

	result, errs = a.generateToken(ctx, tx, data.Platform, data.FcmToken, userData, permissionData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, utils.ErrorInternalServer(err.Error())
	}

	return result, nil
}
