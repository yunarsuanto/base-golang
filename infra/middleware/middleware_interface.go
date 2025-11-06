package middleware

import (
	"net/http"

	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
)

type MiddlewareInterface interface {
	GeneralAccessToken(handlerFunc http.Handler) http.Handler
	// MobileAccessToken(handlerFunc http.Handler) http.Handler
	// WebAccessToken(handlerFunc http.Handler) http.Handler
}

func AccessTokenMiddleware(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) MiddlewareInterface {
	return &accessTokenMiddleware{
		repoCtx,
		infraCtx,
	}
}
