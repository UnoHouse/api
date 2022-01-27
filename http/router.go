package http

import (
	"net/http"

	"github.com/UnoHouse/api/swagger"
	"github.com/gorilla/mux"
)

func NewRouter(h Handler) http.Handler {
	r := mux.NewRouter()

	setupHealthCheck(r)
	setupSwagger(r)

	return r
}

func setupHealthCheck(r *mux.Router) {
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Up"))
	}).Methods(http.MethodGet)
}

func setupSwagger(r *mux.Router) {
	h := http.FileServer(http.FS(swagger.GetStaticFiles()))
	r.PathPrefix("/").Handler(h)
}
