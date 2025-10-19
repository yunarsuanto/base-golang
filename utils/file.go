package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	"github.com/nfnt/resize"
	"github.com/yunarsuanto/base-go/constants"
)

func ResizeImage(data []byte, maxHeight uint) ([]byte, *constants.ErrorResponse) {
	mimeType := http.DetectContentType(data)
	if !InArrayExist(mimeType, constants.JpegMimeType()) {
		return data, nil
	}

	if maxHeight == 0 {
		maxHeight = constants.DefaultMaxImageHeight
	}

	var result []byte
	image, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return result, ErrorInternalServer(err.Error())
	}

	newImage := resize.Resize(0, maxHeight, image, resize.Lanczos3)

	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err = jpeg.Encode(w, newImage, nil)
	if err != nil {
		return result, ErrorInternalServer(err.Error())
	}
	result = b.Bytes()

	return result, nil
}

func CreateFileFromBytes(data []byte, fileName string) (string, string, *constants.ErrorResponse) {
	var result string

	mimeType := mimetype.Detect(data)

	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		return result, mimeType.String(), ErrorInternalServer(err.Error())
	}

	file, err := os.Open(fileName)
	if err != nil {
		return result, mimeType.String(), ErrorInternalServer(err.Error())
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return result, mimeType.String(), ErrorInternalServer(err.Error())
	}

	result, err = filepath.Abs(fileName)
	if err != nil {
		return result, mimeType.String(), ErrorInternalServer(err.Error())
	}

	newPath := fmt.Sprintf("%s%s", result, mimeType.Extension())
	os.Rename(result, newPath)
	result = newPath

	return result, mimeType.String(), nil
}

func GetFileSize(path string) (int64, *constants.ErrorResponse) {
	var result int64
	fileStat, err := os.Stat(path)
	if err != nil {
		return result, ErrorInternalServer(err.Error())
	}
	return fileStat.Size(), nil
}
