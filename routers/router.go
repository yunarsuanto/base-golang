package routers

import (
	"fmt"
	"net/http"

	"github.com/yunarsuanto/base-go/config"
	"github.com/yunarsuanto/base-go/infra/initiator/handler"
	"github.com/yunarsuanto/base-go/infra/middleware"
	v1 "github.com/yunarsuanto/base-go/routers/v1"

	"github.com/gorilla/mux"
)

func InitRouter(handlerCtx *handler.HandlerCtx, mw middleware.MiddlewareInterface, cfg *config.Config) *mux.Router {
	r := mux.NewRouter()

	if cfg.Server.LocalStoragePath != "" {
		staticPath := fmt.Sprintf("/%s", cfg.Server.LocalStoragePath)
		fs := http.FileServer(http.Dir(cfg.Server.LocalStoragePath))
		r.PathPrefix(staticPath).Handler(http.StripPrefix(staticPath, fs))
	}

	s := r.PathPrefix("/api").Subrouter()
	v1.V1Router(handlerCtx, mw, s)

	return r
}
