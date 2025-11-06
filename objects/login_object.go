package objects

import "time"

type LoginRequest struct {
	Username string
	Password string
	Platform string
	FcmToken string
}

type LoginWithGoogleRequest struct {
	Token    string
	Platform string
}

type VerifyRequest struct {
	Token    string
	Platform string
}

type RefreshTokenRequest struct {
	FcmToken string
	Platform string
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
	ExpiredAt    time.Time
	Permissions  []string
	Role         string
}
