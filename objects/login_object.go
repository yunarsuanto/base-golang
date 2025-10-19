package objects

import "time"

type LoginRequest struct {
	Email    string
	Password string
	Platform string
	FcmToken string
}

type RefreshTokenRequest struct {
	FcmToken string
}

type Login struct {
	AccessToken   string
	RefreshToken  string
	ExpiredAt     time.Time
	IsReporter    bool
	IsVerificator bool
	HospitalIds   []string
	Permissions   []string
}

type ForgotPasswordRequest struct {
	Email string
}

type VerifyChangePasswordOtp struct {
	Email string
	Otp   string
}

type ChangePasswordRequest struct {
	Email    string
	Otp      string
	Password string
}
