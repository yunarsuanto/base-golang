package file_handler

import (
	"net/http"

	"github.com/yunarsuanto/base-go/infra/initiator/service"
)

type FileHandlerInterface interface {
	UploadMultipart(w http.ResponseWriter, r *http.Request)
	UploadBase64(w http.ResponseWriter, r *http.Request)
	GetUrl(w http.ResponseWriter, r *http.Request)
}

func NewFileHandler(serviceCtx *service.ServiceCtx) FileHandlerInterface {
	return &handler{
		serviceCtx,
	}
}
