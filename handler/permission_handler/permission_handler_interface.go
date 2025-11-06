package permission_handler

import (
	"net/http"

	"github.com/yunarsuanto/base-go/infra/initiator/service"
)

type PermissionHandlerInterface interface {
	ListPermission(w http.ResponseWriter, r *http.Request)
	CreatePermission(w http.ResponseWriter, r *http.Request)
	UpdatePermission(w http.ResponseWriter, r *http.Request)
	DeletePermission(w http.ResponseWriter, r *http.Request)
}

func NewPermissionHandler(serviceCtx *service.ServiceCtx) PermissionHandlerInterface {
	return &handler{
		serviceCtx,
	}
}
