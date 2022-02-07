package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupOtherRoutes(r *mux.Router, h Handler) {
	r.Methods(http.MethodGet).Path("/app_description").HandlerFunc(BasicAuth(h.getAppDescription()))
	r.Methods(http.MethodPost).Path("/app_description").HandlerFunc(BasicAuth(h.postAppDescription()))
	r.Methods(http.MethodPut).Path("/app_description").HandlerFunc(BasicAuth(h.putAppDescription()))
	r.Methods(http.MethodDelete).Path("/app_description").HandlerFunc(BasicAuth(h.deleteAppDescription()))
	r.Methods(http.MethodPost).Path("/notifications/user/{id}").HandlerFunc(BasicAuth(h.postNotification()))
}
