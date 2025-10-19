package objects

import (
	"net/http"

	common_input_handler "github.com/yunarsuanto/base-go/handler"
)

type CreateActivityLog struct {
	Request      *http.Request
	Module       string
	Body         any
	ResponseMeta common_input_handler.Meta
}
