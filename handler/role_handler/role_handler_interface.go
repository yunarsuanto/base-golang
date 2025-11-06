package role_handler

import (
	"net/http"

	"github.com/yunarsuanto/base-go/infra/initiator/service"
)

type RoleHandlerInterface interface {
	ListRole(w http.ResponseWriter, r *http.Request)
	DetailRole(w http.ResponseWriter, r *http.Request)
	CreateRole(w http.ResponseWriter, r *http.Request)
	UpdateRole(w http.ResponseWriter, r *http.Request)
	DeleteRole(w http.ResponseWriter, r *http.Request)
}

func NewRoleHandler(serviceCtx *service.ServiceCtx) RoleHandlerInterface {
	return &handler{
		serviceCtx,
	}
}
