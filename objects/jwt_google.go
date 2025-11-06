package objects

import "github.com/golang-jwt/jwt/v4"

type GooglePayload struct {
	jwt.StandardClaims
	Aud           string `json:"aud"`
	Azp           string `json:"azp,omitempty"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Exp           int64  `json:"exp"`
	FamilyName    string `json:"family_name,omitempty"`
	GivenName     string `json:"given_name,omitempty"`
	Iat           int64  `json:"iat"`
	Iss           string `json:"iss,omitempty"`
	Jti           string `json:"jti,omitempty"`
	Name          string `json:"name,omitempty"`
	Nbf           int64  `json:"nbf,omitempty"`
	Picture       string `json:"picture,omitempty"`
	Sub           string `json:"sub"`
}
