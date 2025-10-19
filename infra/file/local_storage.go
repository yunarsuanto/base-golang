package file

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/yunarsuanto/base-go/config"
	"github.com/yunarsuanto/base-go/constants"
)

func NewLocalStorage(cfg *config.Config) StorageInterface {
	return &localStorage{
		cfg,
	}
}

type localStorage struct {
	*config.Config
}

func (a localStorage) checkLocalFolderExists() {
	if _, err := os.Stat(a.Server.LocalStoragePath); os.IsNotExist(err) {
		os.Mkdir(a.Server.LocalStoragePath, 0777)
	}
}

func (a localStorage) checkContextFolderExists(context string) {
	contextFolder := fmt.Sprintf("%s/%s", a.Server.LocalStoragePath, context)
	if _, err := os.Stat(contextFolder); os.IsNotExist(err) {
		os.Mkdir(contextFolder, 0777)
	}
}

func (a localStorage) PutFile(data io.Reader, context, path, mimeType string) error {
	a.checkLocalFolderExists()

	pathArr := []string{path}
	if context != "" {
		a.checkContextFolderExists(context)
		pathArr = append([]string{context}, pathArr...)
	}
	fileName := fmt.Sprintf("%s/%s", a.Server.LocalStoragePath, strings.Join(pathArr, "/"))

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)

	err := os.WriteFile(fileName, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (localStorage) DeleteFile(path string) error {
	return errors.New("Unimplemented")
}

func (a localStorage) GetUrl(path string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s", a.Server.AppUrl, a.Server.LocalStoragePath, path)
	return url, nil
}

func (localStorage) GetBytes(path string) ([]byte, error) {
	return []byte{}, errors.New("Unimplemented")
}

func (localStorage) GetProvider() string {
	return constants.LocalStorageProvider
}
