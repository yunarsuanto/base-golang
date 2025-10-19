package jwt

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yunarsuanto/base-go/config"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"

	"github.com/dgrijalva/jwt-go"
)

// JWTInterface interface jwt
type JWTInterface interface {
	ExtractJWTClaims(ctx context.Context, token string, appName, middlewareType string) (claims *objects.JWTClaims, errs *constants.ErrorResponse)
	ValidateTokenIssuer(claims *objects.JWTClaims) *constants.ErrorResponse
	ValidateTokenExpire(ctx context.Context, claims *objects.JWTClaims, reqToken string, appName string, purpose string) *constants.ErrorResponse

	DeleteTokenFromRedis(ctx context.Context, claims *objects.JWTClaims, purpose string, appName string) *constants.ErrorResponse
	GenerateJWTToken(ctx context.Context, request objects.JWTRequest, expireTime time.Duration, appName string, purpose string) (string, *constants.ErrorResponse)
}

// jwtObj struct
type jwtObj struct {
	config             *config.JwtConfig
	redis              *redis.Client
	allowMultipleLogin bool
}

// NewJWT function to connect jwtObj to JWTInterface
// Params:
// cfg: config
// redis: redis client
// Returns JWTInterface
func NewJWT(cfg *config.JwtConfig, redis *redis.Client, allowMultipleLogin bool) JWTInterface {
	return &jwtObj{
		config:             cfg,
		redis:              redis,
		allowMultipleLogin: allowMultipleLogin,
	}
}

// ExtractJWTClaims function to extract jwt claims from authorization header
// Params:
// ctx: context
// token: token to extract from
func (j *jwtObj) ExtractJWTClaims(ctx context.Context, token string, appName, middlewareType string) (claims *objects.JWTClaims, errs *constants.ErrorResponse) {
	// check authorization
	splitToken := strings.Split(token, constants.Bearer)
	if len(splitToken) != 2 {
		return nil, constants.ErrTokenIsRequired
	}
	reqToken := strings.TrimSpace(splitToken[1])

	t, err := jwt.ParseWithClaims(reqToken, &objects.JWTClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(j.config.Secret), nil
	})
	if err != nil {
		return nil, utils.ErrJwt(err)
	}

	claims = t.Claims.(*objects.JWTClaims)
	// Validate Issuer Token
	errs = j.ValidateTokenIssuer(claims)
	if errs != nil {
		return nil, errs
	}

	// Validate token expire
	errs = j.ValidateTokenExpire(ctx, claims, reqToken, appName, claims.Purpose)
	if errs != nil {
		return nil, errs
	}

	return claims, nil
}

// ValidateTokenIssuer is for validate token issuer
func (j *jwtObj) ValidateTokenIssuer(claims *objects.JWTClaims) *constants.ErrorResponse {
	if claims.Issuer != j.config.Issuer {
		return constants.ErrTokenInvalid
	}
	return nil
}

// ValidateTokenExpire is for validate Token Expire
func (j *jwtObj) ValidateTokenExpire(ctx context.Context, claims *objects.JWTClaims, reqToken string, appName string, purpose string) *constants.ErrorResponse {
	redisKey := utils.GenerateRedisKey(appName, claims.Id, purpose, claims.Platform, claims.UniqueKey, j.allowMultipleLogin)
	// check token to redis
	token, err := j.getTokenFromRedis(ctx, redisKey)
	if err != nil {
		return utils.ErrRedis(err.Error(), redisKey)
	}
	if token != reqToken {
		return constants.ErrTokenReplaced
	}

	return nil
}

func (j *jwtObj) getTokenFromRedis(ctx context.Context, key string) (string, error) {
	val, err := j.redis.Get(ctx, key).Result()
	if err != nil {
		return val, err
	}
	return val, nil
}

// DeleteTokenFromRedis function to delete token from redis
// Params:
// ctx: context
// id: user ID / admin ID
// authKey: redis authorization key
// Returns *constants.ErrorResponse
func (j *jwtObj) DeleteTokenFromRedis(ctx context.Context, claims *objects.JWTClaims, purpose string, appName string) *constants.ErrorResponse {
	redisKey := utils.GenerateRedisKey(appName, claims.Id, purpose, claims.Platform, claims.UniqueKey, j.allowMultipleLogin)
	_, err := j.redis.Del(ctx, redisKey).Result()
	if err != nil {
		return utils.ErrRedis(err.Error(), redisKey)
	}

	return nil
}

// Generate Token
// Params:
// ctx: context
// id: user ID / admin ID
// authKey: redis authorization key
// Returns *constants.ErrorResponse
func (j *jwtObj) GenerateJWTToken(ctx context.Context, request objects.JWTRequest, expireTime time.Duration, appName string, purpose string) (string, *constants.ErrorResponse) {
	JWTSignatureKey := []byte(j.config.Secret)
	claims := objects.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        request.Id,
			Issuer:    j.config.Issuer,
			ExpiresAt: time.Now().Add(time.Duration(expireTime)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:          request.Id,
		Purpose:     purpose,
		Platform:    request.Platform,
		UniqueKey:   utils.RandomString(6),
		Permissions: request.Permissions,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	// create token client
	signedToken, err := token.SignedString(JWTSignatureKey)
	if err != nil {
		return "", constants.ErrGenerateToken
	}

	redisKey := utils.GenerateRedisKey(appName, claims.Id, purpose, claims.Platform, claims.UniqueKey, j.allowMultipleLogin)
	errs := j.setTokenToRedis(ctx, signedToken, expireTime, redisKey)
	if errs != nil {
		return "", constants.ErrGenerateToken
	}

	return signedToken, nil
}

func (j *jwtObj) setTokenToRedis(ctx context.Context, token string, expireTime time.Duration, key string) *constants.ErrorResponse {
	_, err := j.redis.Set(ctx, key, token, expireTime).Result()
	if err != nil {
		return constants.ErrSetTokenToRedis
	}

	return nil
}
