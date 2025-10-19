package v1

import (
	"github.com/gorilla/mux"
	"github.com/yunarsuanto/base-go/infra/initiator/handler"
	"github.com/yunarsuanto/base-go/infra/middleware"
)

func generalRouter(handlerCtx *handler.HandlerCtx, mw middleware.MiddlewareInterface, r *mux.Router) {
	a := r.PathPrefix("/general").Subrouter()
	a.Use(mw.GeneralAccessToken)

	// a.HandleFunc("/upload-file", handlerCtx.GeneralFileHandler.UploadMultipart).Methods(http.MethodPost)

	// a.HandleFunc("/location/province/get", handlerCtx.GeneralLocationHandler.GetListProvince).Methods(http.MethodGet)
	// a.HandleFunc(utils.ParsePath("/notification/%s/detail", constants.IdPathVariable), handlerCtx.GeneralNotificationHandler.GetDetail).Methods(http.MethodGet)
}
