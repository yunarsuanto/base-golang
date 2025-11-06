package file_service

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/objects"
	"github.com/yunarsuanto/base-go/utils"
)

func checkFolderTempExists() {
	if _, err := os.Stat("temp/"); os.IsNotExist(err) {
		os.Mkdir("temp/", 0777)
	}
}

func (a service) storeFile(fileContext string, tempFilePath string, mimeType string) (objects.FileUpload, *constants.ErrorResponse) {
	var result objects.FileUpload

	tempFile, err := os.Open(tempFilePath)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}
	defer tempFile.Close()
	defer os.Remove(tempFilePath)

	pathArr := []string{filepath.Base(tempFilePath)}
	if fileContext != "" {
		pathArr = append([]string{fileContext}, pathArr...)
	}
	fullPath := strings.Join(pathArr, "/")

	fileUrl, err := a.Storage.PutFile(tempFile, fileContext, filepath.Base(tempFilePath), a.Config.Server.LocalStoragePath, mimeType)
	if err != nil {
		return result, utils.ErrorInternalServer(err.Error())
	}

	// fileUrl, err := a.Storage.PutFile(tempFile, fileContext, filepath.Base(tempFilePath), a.Config.Server.StorageProvider, mimeType)
	// if err != nil {
	// 	return result, utils.ErrorInternalServer(err.Error())
	// }

	fileSize, errs := utils.GetFileSize(tempFilePath)
	if errs != nil {
		return result, errs
	}

	result = objects.FileUpload{
		Url:      fileUrl,
		Path:     fullPath,
		MimeType: mimeType,
		Size:     fileSize,
	}

	return result, nil
}
