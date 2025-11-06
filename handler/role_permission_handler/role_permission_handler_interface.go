package role_permission_handler

import (
	"net/http"

	"github.com/yunarsuanto/base-go/infra/initiator/service"
)

type RolePermissionHandlerInterface interface {
	UpsertRolePermission(w http.ResponseWriter, r *http.Request)
	DeleteRolePermission(w http.ResponseWriter, r *http.Request)
}

func NewRolePermissionHandler(serviceCtx *service.ServiceCtx) RolePermissionHandlerInterface {
	return &handler{
		serviceCtx,
	}
}
