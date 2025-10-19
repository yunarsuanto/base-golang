package auth_handler

import common_input_handler "github.com/yunarsuanto/base-go/handler"

type loginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,alphanum,min=8" log:"masked"`
	Platform string `json:"platform" validate:"required,oneof=web mobile"`
	FcmToken string `json:"fcmToken"`
}

type refreshTokenRequest struct {
	FcmToken string `json:"fcmToken"`
}

type loginResponseData struct {
	AccessToken   string   `json:"accessToken"`
	RefreshToken  string   `json:"refreshToken"`
	ExpiredAt     string   `json:"expiredAt"`
	IsReporter    bool     `json:"isReporter"`
	IsVerificator bool     `json:"isVerificator"`
	HospitalIds   []string `json:"hospitalIds"`
	Permissions   []string `json:"permissions"`
}

type loginResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *loginResponseData        `json:"data"`
}

type forgotPasswordRequest struct {
	Email string `json:"email" validate:"required"`
}

type forgotPasswordResponseData struct{}

type forgotPasswordResponse struct {
	Meta common_input_handler.Meta   `json:"meta"`
	Data *forgotPasswordResponseData `json:"data"`
}

type verifyChangePasswordOtpRequest struct {
	Email string `json:"email" validate:"required"`
	Otp   string `json:"otp" validate:"required" log:"masked"`
}

type verifyChangePasswordOtpResponseData struct{}

type verifyChangePasswordOtpResponse struct {
	Meta common_input_handler.Meta            `json:"meta"`
	Data *verifyChangePasswordOtpResponseData `json:"data"`
}

type changePasswordRequest struct {
	Email    string `json:"email" validate:"required"`
	Otp      string `json:"otp" validate:"required" log:"masked"`
	Password string `json:"password" validate:"required,alphanum,min=8" log:"masked"`
}

type changePasswordResponseData struct{}

type changePasswordResponse struct {
	Meta common_input_handler.Meta   `json:"meta"`
	Data *changePasswordResponseData `json:"data"`
}

type updatePasswordRequest struct {
	Password string `json:"password" validate:"required,alphanum,min=8" log:"masked"`
}

type updatePasswordResponseData struct{}

type updatePasswordResponse struct {
	Meta common_input_handler.Meta   `json:"meta"`
	Data *updatePasswordResponseData `json:"data"`
}

type logoutResponseData struct{}

type logoutResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *logoutResponseData       `json:"data"`
}
