package auth_service

import (
	"context"
	"fmt"
	"html/template"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/infra/jwt"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type service struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (a service) Login(ctx context.Context, data objects.LoginRequest) (objects.LoginResponse, *constants.ErrorResponse) {
	var result objects.LoginResponse
	isSuperAdmin := false

	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	userData, errs := a.UserRepo.GetByUsername(ctx, tx, data.Username)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}
	if !userData.IsActive {
		_ = tx.Rollback()
		return result, constants.ErrInactiveUser
	}

	if userData.Username == "superadmin" {
		isSuperAdmin = true
	}

	roleName, permissionData, errs := a.getAdminPermissions(ctx, tx, userData.Id, isSuperAdmin)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	isCheckPassword := utils.CheckPasswordHash(data.Password, userData.Password)
	if !isCheckPassword {
		_ = tx.Rollback()
		return result, constants.ErrEmailAndPasswordNotMatch
	}

	if userData.Username == "superadmin" {
		roleName = "superadmin"
	}

	result, errs = a.generateToken(ctx, data.Platform, data.FcmToken, userData, permissionData, isSuperAdmin, roleName)
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

func (a service) LoginWithGoogle(ctx context.Context, data objects.LoginWithGoogleRequest) (objects.LoginResponse, *constants.ErrorResponse) {
	var result objects.LoginResponse
	isSuperAdmin := false
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	claims, err := jwt.VerifyGoogleJWT(data.Token, a.Config.JwtConfig.ClientId)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}
	if !claims.EmailVerified {
		return result, constants.ErrInactiveGoogleAccount
	}

	userData, errs := a.UserRepo.GetByUsername(ctx, tx, claims.Email)
	if errs != nil {
		if errs.HttpCode != 404 {
			_ = tx.Rollback()
			return result, errs
		}
	}

	if userData.Id == "" {
		pass, errHash := utils.HashPassword(utils.RandomString(10))
		if errHash != nil {
			return result, errs
		}
		errs = a.UserRepo.CreateUser(ctx, tx, models.CreateUser{
			Username:   claims.Email,
			Password:   pass,
			IsActive:   false,
			ProviderId: utils.ConvertStringToNil(data.Token),
			Provider:   utils.ConvertStringToNil("google"),
		})
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		userData, errs = a.UserRepo.GetByUsername(ctx, tx, claims.Email)
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		errs = a.UserRepo.CreateUserProfile(ctx, tx, models.CreateUserProfile{
			UserId:          userData.Id,
			NationalId:      "",
			Fullname:        fmt.Sprintf("%s %s", claims.GivenName, claims.FamilyName),
			Email:           claims.Email,
			Phone:           "",
			Address:         "",
			PostalCode:      "",
			Age:             0,
			Latitude:        0,
			Longitude:       0,
			ProfileImage:    claims.Picture,
			NationalIdImage: "",
			GuardianName:    "",
		})
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	if !userData.IsActive {
		randomString := utils.RandomString(10)
		errs = a.UserRepo.UpdateTokenVerification(ctx, tx, models.UpdateUserTokenVerification{
			Username: claims.Email,
			Token:    randomString,
		})
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
		errs = a.Mailer.SendCustom(objects.SendCustomMailData{
			To:      []string{claims.Email},
			Title:   "verification mail",
			Content: template.HTML(fmt.Sprintf(`klik link tersebur <a href="%s">klik untuk verifikasi</a> `, a.Config.Server.UrlFrontend+"/"+randomString)),
		})
		if errs != nil {
			return result, constants.Error(403, "email tidak terkirim")
		}

		if tx == nil {
			return result, utils.ErrorInternalServer("transaction is nil")
		}
		err = tx.Commit()
		if err != nil {
			return result, utils.ErrorInternalServer(err.Error())
		}
		return result, constants.ErrVerifyActivation
	}

	if userData.Username == "superadmin" {
		isSuperAdmin = true
	}

	roleName, permissionData, errs := a.getAdminPermissions(ctx, tx, userData.Id, isSuperAdmin)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	claimsInDB, err := jwt.VerifyGoogleJWT(utils.NullScan(userData.ProviderId), a.Config.JwtConfig.ClientId)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	if claimsInDB.Email != claims.Email {
		return result, constants.ErrEmailAndPasswordNotMatch
	} else {
		errs = a.UserRepo.UpdateUser(ctx, tx, models.UpdateUser{
			Id:         userData.Id,
			Username:   userData.Username,
			ProviderId: utils.ToNullScan(data.Token, ""),
		})
		if errs != nil {
			_ = tx.Rollback()
			return result, errs
		}
	}

	if userData.Username == "superadmin" {
		roleName = "superadmin"
	}

	result, errs = a.generateToken(ctx, data.Platform, "", userData, permissionData, isSuperAdmin, roleName)
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

func (a service) Verify(ctx context.Context, data objects.VerifyRequest) (objects.LoginResponse, *constants.ErrorResponse) {
	var result objects.LoginResponse
	var roles []models.ListRole
	isSuperAdmin := false
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	userData, errs := a.UserRepo.GetByTokenVerification(ctx, tx, data.Token)
	if errs != nil {
		if errs.HttpCode != 404 {
			_ = tx.Rollback()
			return result, errs
		}
	}

	errs = utils.CompareToken(data.Token, utils.NullScan(userData.TokenVerification))
	if errs != nil {
		_ = tx.Rollback()
		return result, constants.ErrEmailAndPasswordNotMatch
	}

	errs = a.UserRepo.UpdateTokenVerificationIsActiveUser(ctx, tx, models.UpdateUserIsActiveTokenVerification{
		Username: userData.Username,
	})
	if errs != nil {
		_ = tx.Rollback()
		return result, constants.ErrPasswordNotMatch
	}

	if userData.Username == "superadmin" {
		isSuperAdmin = true
	}

	roles, errs = a.RoleRepo.ListRole(ctx, tx, &objects.Pagination{
		Limit:  1,
		Page:   1,
		Search: "user",
	})
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	errs = a.UserRoleRepo.UpsertUserRole(ctx, tx, models.UpsertUserRoleRequest{
		UserId:   userData.Id,
		RoleId:   roles[0].Id,
		IsActive: true,
	})
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	roleName, permissionData, errs := a.getAdminPermissions(ctx, tx, userData.Id, isSuperAdmin)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result, errs = a.generateToken(ctx, data.Platform, "", userData, permissionData, false, roleName)
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

func (a service) RefreshToken(ctx context.Context, data objects.RefreshTokenRequest) (objects.LoginResponse, *constants.ErrorResponse) {
	var result objects.LoginResponse
	var roles []models.ListRole
	isSuperAdmin := false
	tx, err := a.Db.Begin(ctx)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	claims, errs := a.Jwt.ExtractJwtRefrehToken(ctx, data.FcmToken)
	if errs != nil {
		return result, constants.ErrTokenInvalid
	}

	userData, errs := a.UserRepo.GetById(ctx, tx, claims.Id)
	if errs != nil {
		if errs.HttpCode != 404 {
			_ = tx.Rollback()
			return result, errs
		}
	}

	if userData.Username == "superadmin" {
		isSuperAdmin = true
	}

	roles, errs = a.RoleRepo.ListRole(ctx, tx, &objects.Pagination{
		Limit:  1,
		Page:   1,
		Search: "user",
	})
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	errs = a.UserRoleRepo.UpsertUserRole(ctx, tx, models.UpsertUserRoleRequest{
		UserId:   userData.Id,
		RoleId:   roles[0].Id,
		IsActive: true,
	})
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	roleName, permissionData, errs := a.getAdminPermissions(ctx, tx, userData.Id, isSuperAdmin)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	if userData.Username == "superadmin" {
		roleName = "superadmin"
	}

	result, errs = a.generateToken(ctx, data.Platform, "", userData, permissionData, isSuperAdmin, roleName)
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
