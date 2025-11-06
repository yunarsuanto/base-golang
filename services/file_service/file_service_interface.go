package file_service

import (
	"context"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
)

type FileServiceInterface interface {
	UploadMultipart(ctx context.Context, data objects.FileUploadMultipartRequest) (objects.FileUpload, *constants.ErrorResponse)
	UploadBase64(ctx context.Context, data objects.FileUploadBase64Request) (objects.FileUpload, *constants.ErrorResponse)
	GetUrl(ctx context.Context, data objects.FileGetUrlRequest) (objects.FileGetUrl, *constants.ErrorResponse)
}

func NewFileService(repoCtx *repository.RepoCtx, infraCtx *infra.InfraCtx) FileServiceInterface {
	return &service{
		repoCtx,
		infraCtx,
	}
}
