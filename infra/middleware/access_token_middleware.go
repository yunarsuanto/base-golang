package middleware

import (
	"context"
	"net/http"

	"github.com/yunarsuanto/base-go/constants"
	common_input_handler "github.com/yunarsuanto/base-go/handler"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type accessTokenMiddleware struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (a accessTokenMiddleware) extractAccessToken(ctx context.Context, r *http.Request, middlewareType string) (*objects.JWTClaims, *constants.ErrorResponse) {
	var claims *objects.JWTClaims

	tokenHeader := r.Header.Get(constants.HeaderAuthorization)
	if tokenHeader == "" {
		return claims, constants.ErrIneligibleAccess
	}

	claims, errs := a.Jwt.ExtractJWTClaims(ctx, tokenHeader, constants.AppName, middlewareType)
	if errs != nil {
		return claims, errs
	}

	return claims, nil
}

func (a accessTokenMiddleware) GeneralAccessToken(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		claims, errs := a.extractAccessToken(ctx, r, "")
		if errs != nil {
			result := common_input_handler.ErrorResponse{Meta: utils.SetErrorMeta(errs)}
			utils.JSONResponse(w, errs.HttpCode, &result)
			return
		}

		ctx = context.WithValue(ctx, constants.ClaimsContextKey, claims)
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

// func (a accessTokenMiddleware) MobileAccessToken(handlerFunc http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()

// 		claims, errs := a.extractAccessToken(ctx, r, "")
// 		if errs != nil {
// 			result := common_input_handler.ErrorResponse{Meta: utils.SetErrorMeta(errs)}
// 			utils.JSONResponse(w, errs.HttpCode, &result)
// 			return
// 		}

// 		// tx, err := a.Db.Begin(ctx)
// 		// if err != nil {
// 		// 	errs = utils.ErrorInternalServer(err.Error())
// 		// 	result := common_input_handler.ErrorResponse{Meta: utils.SetErrorMeta(errs)}
// 		// 	utils.JSONResponse(w, errs.HttpCode, &result)
// 		// 	return
// 		// }

// 		// hospitalId := r.Header.Get(constants.HeaderHospitalId)
// 		// errs = a.UsersHospitalsRepo.ValidateUserIdHospitalId(ctx, tx, claims.Id, hospitalId)
// 		// if errs != nil {
// 		// 	_ = tx.Rollback()
// 		// 	result := common_input_handler.ErrorResponse{Meta: utils.SetErrorMeta(errs)}
// 		// 	utils.JSONResponse(w, errs.HttpCode, &result)
// 		// 	return
// 		// }
// 		// claims.HospitalId = hospitalId

// 		// err = tx.Commit()
// 		// if err != nil {
// 		// 	errs = utils.ErrorInternalServer(err.Error())
// 		// 	result := common_input_handler.ErrorResponse{Meta: utils.SetErrorMeta(errs)}
// 		// 	utils.JSONResponse(w, errs.HttpCode, &result)
// 		// 	return
// 		// }

// 		if claims.Platform != constants.MobilePlatform {
// 			errs = constants.ErrIneligibleAccess
// 			result := common_input_handler.ErrorResponse{Meta: utils.SetErrorMeta(errs)}
// 			utils.JSONResponse(w, errs.HttpCode, &result)
// 			return
// 		}

// 		hasPermission := utils.InArrayExist(constants.PermissionAccessMobile, claims.Permissions)
// 		if !hasPermission {
// 			errs = constants.ErrIneligibleAccess
// 			result := common_input_handler.ErrorResponse{Meta: utils.SetErrorMeta(errs)}
// 			utils.JSONResponse(w, errs.HttpCode, &result)
// 			return
// 		}

// 		ctx = context.WithValue(ctx, constants.ClaimsContextKey, claims)
// 		r = r.WithContext(ctx)
// 		handlerFunc.ServeHTTP(w, r)
// 	})
// }

// func (a accessTokenMiddleware) WebAccessToken(handlerFunc http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()

// 		claims, errs := a.extractAccessToken(ctx, r, "")
// 		if errs != nil {
// 			result := common_input_handler.ErrorResponse{Meta: utils.SetErrorMeta(errs)}
// 			utils.JSONResponse(w, errs.HttpCode, &result)
// 			return
// 		}

// 		if claims.Platform != constants.WebPlatform {
// 			errs = constants.ErrIneligibleAccess
// 			result := common_input_handler.ErrorResponse{Meta: utils.SetErrorMeta(errs)}
// 			utils.JSONResponse(w, errs.HttpCode, &result)
// 			return
// 		}

// 		hasPermission := utils.InArrayExist(constants.PermissionAccessWeb, claims.Permissions)
// 		if !hasPermission {
// 			errs = constants.ErrIneligibleAccess
// 			result := common_input_handler.ErrorResponse{Meta: utils.SetErrorMeta(errs)}
// 			utils.JSONResponse(w, errs.HttpCode, &result)
// 			return
// 		}

// 		ctx = context.WithValue(ctx, constants.ClaimsContextKey, claims)
// 		r = r.WithContext(ctx)
// 		handlerFunc.ServeHTTP(w, r)
// 	})
// }
