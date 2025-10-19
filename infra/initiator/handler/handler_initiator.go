package handler

import (
	"github.com/yunarsuanto/base-go/handler/auth_handler"
	"github.com/yunarsuanto/base-go/infra/initiator/service"
)

type HandlerCtx struct {
	AuthHandler auth_handler.AuthHandlerInterface
}

func InitHandlerCtx(serviceCtx *service.ServiceCtx) *HandlerCtx {
	return &HandlerCtx{
		AuthHandler: auth_handler.NewAuthHandler(serviceCtx),
	}
}
