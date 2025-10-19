package file

import (
	"fmt"
	"io"

	"github.com/yunarsuanto/base-go/config"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/utils"
)

type FileCtx struct {
	Config   config.ServerConfig
	Firebase StorageInterface
	Local    StorageInterface
	Url      StorageInterface
}

func (f FileCtx) GetUrl(path, pathType *string) (string, *constants.ErrorResponse) {
	var result string

	filePath := utils.NullScan(path)
	storageProvider := utils.NullScan(pathType)

	if filePath == "" {
		return result, nil
	}
	if storageProvider == "" {
		storageProvider = f.Config.StorageProvider
	}

	var err error
	switch storageProvider {
	case constants.FirebaseStorageProvider:
		result, err = f.Firebase.GetUrl(filePath)
		if err != nil {
			return result, utils.ErrorInternalServer(err.Error())
		}
	case constants.LocalStorageProvider:
		result, err = f.Local.GetUrl(filePath)
		if err != nil {
			return result, utils.ErrorInternalServer(err.Error())
		}
	case constants.UrlStorageProvider:
		result, err = f.Url.GetUrl(filePath)
		if err != nil {
			return result, utils.ErrorInternalServer(err.Error())
		}
	}

	return result, nil
}

func (f FileCtx) GetBytes(path, pathType string) ([]byte, *constants.ErrorResponse) {
	var result []byte
	var err error
	switch pathType {
	case constants.FirebaseStorageProvider:
		result, err = f.Firebase.GetBytes(path)
		if err != nil {
			return result, utils.ErrorInternalServer(err.Error())
		}
	case constants.LocalStorageProvider:
		result, err = f.Local.GetBytes(path)
		if err != nil {
			return result, utils.ErrorInternalServer(err.Error())
		}
	case constants.UrlStorageProvider:
		result, err = f.Url.GetBytes(path)
		if err != nil {
			return result, utils.ErrorInternalServer(err.Error())
		}
	}

	return result, nil
}

func (f FileCtx) PutFile(data io.Reader, context, path, pathType, mimeType string) (string, error) {
	var result string
	switch pathType {
	case constants.FirebaseStorageProvider:
		err := f.Firebase.PutFile(data, context, path, mimeType)
		if err != nil {
			return result, err
		}
		result, err = f.Firebase.GetUrl(path)
		if err != nil {
			return result, err
		}
	case constants.LocalStorageProvider:
		err := f.Local.PutFile(data, context, path, mimeType)
		if err != nil {
			return result, err
		}
		fullPath := path
		if context != "" {
			fullPath = fmt.Sprintf("%s/%s", context, path)
		}
		result, err = f.Local.GetUrl(fullPath)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}
