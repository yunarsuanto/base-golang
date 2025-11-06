package jwt

import (
	"fmt"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yunarsuanto/base-go/objects"
)

const googleJWKSURL = "https://www.googleapis.com/oauth2/v3/certs"

func VerifyGoogleJWT(tokenString string, clientID string) (*objects.GooglePayload, error) {
	// Load JWKS
	jwks, err := keyfunc.Get(googleJWKSURL, keyfunc.Options{
		RefreshInterval: time.Hour,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to load JWKS: %v", err)
	}

	// Parse & verify
	token, err := jwt.ParseWithClaims(tokenString, &objects.GooglePayload{}, jwks.Keyfunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse/verify JWT: %v", err)
	}

	// Pastikan valid
	if claims, ok := token.Claims.(*objects.GooglePayload); ok && token.Valid {
		// Optional: cek aud dan iss
		if claims.Aud != clientID {
			return nil, fmt.Errorf("invalid aud claim")
		}
		if claims.Iss != "https://accounts.google.com" && claims.Iss != "accounts.google.com" {
			return nil, fmt.Errorf("invalid iss claim")
		}

		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}
