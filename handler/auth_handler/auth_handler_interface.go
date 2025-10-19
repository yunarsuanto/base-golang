package auth_handler

import (
	"net/http"

	"github.com/yunarsuanto/base-go/infra/initiator/service"
)

type AuthHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler(serviceCtx *service.ServiceCtx) AuthHandlerInterface {
	return &handler{
		serviceCtx,
	}
}
