package file_service

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"time"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/infra/initiator/infra"
	"github.com/yunarsuanto/base-go/infra/initiator/repository"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

type service struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (a service) UploadMultipart(ctx context.Context, data objects.FileUploadMultipartRequest) (objects.FileUpload, *constants.ErrorResponse) {
	var result objects.FileUpload

	fileName := fmt.Sprintf("temp/%s-%s", utils.RandomString(40), time.Now().Format("200601021504050700"))
	checkFolderTempExists()

	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, data.File)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}
	fileByte := buf.Bytes()

	var errs *constants.ErrorResponse
	fileByte, errs = utils.ResizeImage(fileByte, uint(a.Config.Server.MaxImageHeight))
	if errs != nil {
		return result, errs
	}

	tempFilePath, mimeType, errs := utils.CreateFileFromBytes(fileByte, fileName)
	if errs != nil {
		return result, errs
	}

	if !utils.InArrayExist(mimeType, constants.ValidFileMimeType()) {
		return result, constants.ErrInvalidMimeType
	}

	result, errs = a.storeFile(data.Context, tempFilePath, mimeType)
	if errs != nil {
		return result, errs
	}

	return result, nil
}

func (a service) UploadBase64(ctx context.Context, data objects.FileUploadBase64Request) (objects.FileUpload, *constants.ErrorResponse) {
	var result objects.FileUpload

	fileName := fmt.Sprintf("temp/%s-%s", utils.RandomString(40), time.Now().Format("200601021504050700"))
	checkFolderTempExists()

	fileByte, err := base64.StdEncoding.DecodeString(data.File)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	var errs *constants.ErrorResponse
	fileByte, errs = utils.ResizeImage(fileByte, uint(a.Config.Server.MaxImageHeight))
	if errs != nil {
		return result, errs
	}

	tempFilePath, mimeType, errs := utils.CreateFileFromBytes(fileByte, fileName)
	if errs != nil {
		return result, errs
	}

	if !utils.InArrayExist(mimeType, constants.ValidFileMimeType()) {
		return result, constants.ErrInvalidMimeType
	}

	result, errs = a.storeFile(data.Context, tempFilePath, mimeType)
	if errs != nil {
		return result, errs
	}

	return result, nil
}

func (a service) GetUrl(ctx context.Context, data objects.FileGetUrlRequest) (objects.FileGetUrl, *constants.ErrorResponse) {
	var result objects.FileGetUrl

	fileUrl, errs := a.Storage.GetUrl(&data.Path, nil)
	if errs != nil {
		return result, errs
	}

	result = objects.FileGetUrl{
		Url: fileUrl,
	}

	return result, nil
}
