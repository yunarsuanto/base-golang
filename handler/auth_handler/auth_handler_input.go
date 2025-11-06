package auth_handler

import common_input_handler "github.com/yunarsuanto/base-go/handler"

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,alphanum,min=8" log:"masked"`
	Platform string `json:"platform" validate:"required,oneof=web mobile"`
	FcmToken string `json:"fcmToken"`
}

type loginWithGoogleRequest struct {
	Token    string `json:"token" validate:"required"`
	Platform string `json:"platform" validate:"required"`
}

type verifyRequest struct {
	Token    string `json:"token" validate:"required"`
	Platform string `json:"platform" validate:"required"`
}

type refreshTokenRequest struct {
	FcmToken string `json:"fcm_token" validate:"required"`
	Platform string `json:"platform" validate:"required"`
}

type loginResponseData struct {
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	ExpiredAt    string   `json:"expiredAt"`
	Permissions  []string `json:"permissions"`
	Role         string   `json:"role"`
}

type loginResponse struct {
	Meta common_input_handler.Meta `json:"meta"`
	Data *loginResponseData        `json:"data"`
}
