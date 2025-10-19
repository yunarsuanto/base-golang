package objects

import "github.com/dgrijalva/jwt-go"

// JWTClaims struct of standard JWT middleware
type JWTClaims struct {
	jwt.StandardClaims
	Id          string
	Purpose     string
	Platform    string
	UniqueKey   string
	HospitalId  string
	Permissions []string
}

// JWTRequest struct for request jwt
type JWTRequest struct {
	Id          string   `json:"id"`
	Platform    string   `json:"platform"`
	Permissions []string `json:"permissions"`
}

// JWTSimpleRequest struct for request jwt
type JWTSimpleRequest struct {
	Uid    string               `json:"uid"`
	Claims JWTSimpleChildClaims `json:"claims"`
}

// JWTSimpleChildClaims struct for request jwt
type JWTSimpleChildClaims struct {
	Uid string `json:"uid"`
	Alg string `json:"alg"`
}

// JWTSimpleClaims struct for claims response jwt
type JWTSimpleClaims struct {
	jwt.StandardClaims
	JWTSimpleRequest
}
