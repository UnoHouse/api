package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setupUserRoutes(r *mux.Router, h Handler) {
	r.Methods(http.MethodGet).Path("/users/{id}").HandlerFunc(BasicAuth(h.getUsersId()))
	r.Methods(http.MethodPost).Path("/users/{id}/edit").HandlerFunc(BasicAuth(h.postUsersIdEdit()))
	r.Methods(http.MethodGet).Path("/users/{id}/devices").HandlerFunc(BasicAuth(h.getUserDevices()))
	r.Methods(http.MethodPost).Path("/users/{id}/devices").HandlerFunc(BasicAuth(h.postUserDevices()))
	r.Methods(http.MethodGet).Path("/users/{id}/notifications").HandlerFunc(BasicAuth(h.getUsersIdNotifications()))
	r.Methods(http.MethodPost).Path("/users/change_password").HandlerFunc(BasicAuth(h.postUsersIdChangePassword()))
	r.Methods(http.MethodPost).Path("/users/auth").HandlerFunc(h.postUsersLogin())
	r.Methods(http.MethodPost).Path("/users/registration").HandlerFunc(BasicAuth(h.postUsersRegistration()))
	r.Methods(http.MethodPost).Path("/users/exist").HandlerFunc(BasicAuth(h.postUsersExist()))
}
