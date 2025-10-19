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
		Email:    in.Email,
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
			AccessToken:   resultData.AccessToken,
			RefreshToken:  resultData.RefreshToken,
			ExpiredAt:     resultData.ExpiredAt.Format(time.RFC3339),
			IsReporter:    resultData.IsReporter,
			IsVerificator: resultData.IsVerificator,
			HospitalIds:   resultData.HospitalIds,
			Permissions:   resultData.Permissions,
		},
	}
	utils.JSONResponse(w, result.Meta.Status, &result)
}
