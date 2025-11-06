package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yunarsuanto/base-go/infra/initiator/handler"
	"github.com/yunarsuanto/base-go/infra/middleware"
)

func generalRouter(handlerCtx *handler.HandlerCtx, mw middleware.MiddlewareInterface, r *mux.Router) {
	r.PathPrefix("/general/local/").Handler(
		http.StripPrefix("/api/v1/general/local/", http.FileServer(http.Dir("./local"))),
	)

	a := r.PathPrefix("/").Subrouter()
	a.HandleFunc("/login", handlerCtx.AuthHandler.Login).Methods(http.MethodPost)
	a.HandleFunc("/login/google", handlerCtx.AuthHandler.LoginWithGoogle).Methods(http.MethodPost)
	a.HandleFunc("/verify", handlerCtx.AuthHandler.Verify).Methods(http.MethodPost)
	a.HandleFunc("/refresh-token", handlerCtx.AuthHandler.RefreshToken).Methods(http.MethodPost)

	b := r.PathPrefix("/general").Subrouter()
	b.Use(mw.GeneralAccessToken)
	b.HandleFunc("/upload-file", handlerCtx.FileHandler.UploadBase64).Methods(http.MethodPost)

	// a.HandleFunc("/location/province/get", handlerCtx.GeneralLocationHandler.GetListProvince).Methods(http.MethodGet)
	// a.HandleFunc(utils.ParsePath("/notification/%s/detail", constants.IdPathVariable), handlerCtx.GeneralNotificationHandler.GetDetail).Methods(http.MethodGet)
}
