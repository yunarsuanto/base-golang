package file

import (
	"io"
)

type StorageInterface interface {
	PutFile(data io.Reader, context, path, mimeType string) error
	DeleteFile(path string) error
	GetUrl(path string) (string, error)
	GetBytes(path string) ([]byte, error)
	GetProvider() string
}
