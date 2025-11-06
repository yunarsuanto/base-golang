package user_handler

import (
	"net/http"

	"github.com/yunarsuanto/base-go/infra/initiator/service"
)

type UserHandlerInterface interface {
	ListUser(w http.ResponseWriter, r *http.Request)
	DetailUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(serviceCtx *service.ServiceCtx) UserHandlerInterface {
	return &handler{
		serviceCtx,
	}
}
