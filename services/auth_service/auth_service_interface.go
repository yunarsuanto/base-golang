package auth_service

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, data objects.LoginRequest) (objects.LoginResponse, *constants.ErrorResponse)
	LoginWithGoogle(ctx context.Context, data objects.LoginWithGoogleRequest) (objects.LoginResponse, *constants.ErrorResponse)
	Verify(ctx context.Context, data objects.VerifyRequest) (objects.LoginResponse, *constants.ErrorResponse)
	RefreshToken(ctx context.Context, data objects.RefreshTokenRequest) (objects.LoginResponse, *constants.ErrorResponse)
}

func NewAuthService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) AuthServiceInterface {
	return &service{
		repoCtx,
		infraCtx,
	}
}
