package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/yunarsuanto/base-go/constants"
	common_input_handler "github.com/yunarsuanto/base-go/handler"

	log "github.com/sirupsen/logrus"
)

const (
	contentType              = "Content-Type"
	contentTypeValue         = "application/json; charset=utf-8"
	xContentTypeOptions      = "X-Content-Type-Options"
	xContentTypeOptionsValue = "nosniff"
)

func JSONResponse(w http.ResponseWriter, statusCode int, r any) {
	w.Header().Set(contentType, contentTypeValue)
	w.Header().Set(xContentTypeOptions, xContentTypeOptionsValue)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Error(err)
	}
}

func SetErrorMeta(errs *constants.ErrorResponse, permissions ...string) common_input_handler.Meta {
	PrintError(*errs)
	var permissionCode string
	if len(permissions) > 0 {
		permissionCode = permissions[0]
	}
	return common_input_handler.Meta{
		Message:        errs.Err.Error(),
		Status:         errs.HttpCode,
		PermissionCode: permissionCode,
	}
}

func SetSuccessMeta(message string, permissions ...string) common_input_handler.Meta {
	status := http.StatusOK
	if strings.Contains(message, "Create") {
		status = http.StatusCreated
	}

	var permissionCode string
	if len(permissions) > 0 {
		permissionCode = permissions[0]
	}
	return common_input_handler.Meta{
		Message:        message,
		Status:         status,
		PermissionCode: permissionCode,
	}
}
