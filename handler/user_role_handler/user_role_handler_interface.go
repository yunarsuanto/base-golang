package user_role_handler

import (
	"net/http"

	"github.com/yunarsuanto/base-go/infra/initiator/service"
)

type UserRoleHandlerInterface interface {
	UpsertUserRole(w http.ResponseWriter, r *http.Request)
	DeleteUserRole(w http.ResponseWriter, r *http.Request)
}

func NewUserRoleHandler(serviceCtx *service.ServiceCtx) UserRoleHandlerInterface {
	return &handler{
		serviceCtx,
	}
}
