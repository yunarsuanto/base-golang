package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yunarsuanto/base-go/infra/initiator/handler"
	"github.com/yunarsuanto/base-go/infra/middleware"
)

func V1Router(handlerCtx *handler.HandlerCtx, mw middleware.MiddlewareInterface, r *mux.Router) {
	a := r.PathPrefix("/v1").Subrouter()
	a.HandleFunc("/login", handlerCtx.AuthHandler.Login).Methods(http.MethodPost)

	// a.HandleFunc(utils.ParsePath("/form/%s/version/%s/detail", constants.FormIdPathVariable, constants.IdPathVariable), handlerCtx.GeneralFormVersionHandler.GetDetail).Methods(http.MethodGet)

	generalRouter(handlerCtx, mw, a)
}
