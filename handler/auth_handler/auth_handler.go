package auth_handler

import (
	"net/http"
	"time"

	"github.com/yunarsuanto/base-go/infra/initiator/service"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type handler struct {
	*service.ServiceCtx
}

func (a handler) Login(w http.ResponseWriter, r *http.Request) {
	var result loginResponse

	ctx := r.Context()
	var in loginRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs := utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = loginResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	data := objects.LoginRequest{
		Username: in.Username,
		Password: in.Password,
		Platform: in.Platform,
		FcmToken: in.FcmToken,
	}
	resultData, errs := a.AuthService.Login(ctx, data)
	if errs != nil {
		result = loginResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = loginResponse{
		Meta: utils.SetSuccessMeta("Login"),
		Data: &loginResponseData{
			AccessToken:  resultData.AccessToken,
			RefreshToken: resultData.RefreshToken,
			ExpiredAt:    resultData.ExpiredAt.Format(time.RFC3339),
			Permissions:  resultData.Permissions,
			Role:         resultData.Role,
		},
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) LoginWithGoogle(w http.ResponseWriter, r *http.Request) {
	var result loginResponse

	ctx := r.Context()
	var in loginWithGoogleRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs := utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = loginResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	data := objects.LoginWithGoogleRequest(in)
	resultData, errs := a.AuthService.LoginWithGoogle(ctx, data)
	if errs != nil {
		result = loginResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = loginResponse{
		Meta: utils.SetSuccessMeta("Login"),
		Data: &loginResponseData{
			AccessToken:  resultData.AccessToken,
			RefreshToken: resultData.RefreshToken,
			ExpiredAt:    resultData.ExpiredAt.Format(time.RFC3339),
			Permissions:  resultData.Permissions,
			Role:         resultData.Role,
		},
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) Verify(w http.ResponseWriter, r *http.Request) {
	var result loginResponse

	ctx := r.Context()
	var in verifyRequest
	defer func() {
		a.ActivityLogService.Create(ctx, objects.CreateActivityLog{
			Request:      r,
			Body:         utils.MaskBody(&in),
			ResponseMeta: result.Meta,
		})
	}()

	errs := utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = loginResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	data := objects.VerifyRequest(in)
	resultData, errs := a.AuthService.Verify(ctx, data)
	if errs != nil {
		result = loginResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = loginResponse{
		Meta: utils.SetSuccessMeta("Verify"),
		Data: &loginResponseData{
			AccessToken:  resultData.AccessToken,
			RefreshToken: resultData.RefreshToken,
			ExpiredAt:    resultData.ExpiredAt.Format(time.RFC3339),
			Permissions:  resultData.Permissions,
			Role:         resultData.Role,
		},
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}

func (a handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var result loginResponse

	ctx := r.Context()
	var in refreshTokenRequest
	errs := utils.DecodeJson(&in, r.Body)
	if errs != nil {
		result = loginResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	data := objects.RefreshTokenRequest(in)
	resultData, errs := a.AuthService.RefreshToken(ctx, data)
	if errs != nil {
		result = loginResponse{Meta: utils.SetErrorMeta(errs)}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = loginResponse{
		Meta: utils.SetSuccessMeta("RefreshToken"),
		Data: &loginResponseData{
			AccessToken:  resultData.AccessToken,
			RefreshToken: resultData.RefreshToken,
			ExpiredAt:    resultData.ExpiredAt.Format(time.RFC3339),
			Permissions:  resultData.Permissions,
			Role:         resultData.Role,
		},
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
