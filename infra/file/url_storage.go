package file

import (
	"errors"
	"io"

	"github.com/yunarsuanto/base-go/config"
	"github.com/yunarsuanto/base-go/constants"
)

func NewUrlStorage(cfg *config.Config) StorageInterface {
	return &urlStorage{
		cfg,
	}
}

type urlStorage struct {
	*config.Config
}

func (urlStorage) PutFile(data io.Reader, context, path, mimeType string) error {
	return errors.New("Unimplemented")
}

func (urlStorage) DeleteFile(path string) error {
	return errors.New("Unimplemented")
}

func (urlStorage) GetUrl(path string) (string, error) {
	return path, nil
}

func (urlStorage) GetBytes(path string) ([]byte, error) {
	return []byte{}, errors.New("Unimplemented")
}

func (urlStorage) GetProvider() string {
	return constants.UrlStorageProvider
}
