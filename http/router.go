package http

import (
	"encoding/json"
	"net/http"

	"github.com/UnoHouse/api/swagger"
	"github.com/UnoHouse/api/utils/logger"
	"github.com/gorilla/mux"
)

type AppRoute struct {
	Path    string   `json:"path"`
	Methods []string `json:"methods"`
}

var appRoutes []AppRoute

func NewRouter(handler Handler) http.Handler {
	router := mux.NewRouter()

	setupHealthCheck(router)
	setupSwagger(router)
	setupIndexPageRoutesInfo(router)
	setupUserRoutes(router, handler)
	setupOtherRoutes(router, handler)

	getRoutesList(router)
	return router
}

func SetupRoutes(router *mux.Router, handler Handler) {

}

func setupHealthCheck(r *mux.Router) {
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Up"))
	}).Methods(http.MethodGet)
}

func setupIndexPageRoutesInfo(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		js, err := json.MarshalIndent(appRoutes, "", "    ")
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
		}

		respond(w, http.StatusOK, js)
	}).Methods(http.MethodGet)

}

func setupSwagger(r *mux.Router) {
	h := http.FileServer(http.FS(swagger.GetStaticFiles()))
	r.PathPrefix("/swagger").Handler(h).Methods(http.MethodGet)
}

func getRoutesList(router *mux.Router) []AppRoute {
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		routePath, err := route.GetPathTemplate()
		if err != nil {
			logger.Log(err.Error(), logger.Error)
		}
		routeMethods, err := route.GetMethods()
		if err != nil {
			logger.Log(err.Error(), logger.Error)
		}
		appRoutes = append(appRoutes, AppRoute{Path: routePath, Methods: routeMethods})
		return nil
	})
	return appRoutes
}
