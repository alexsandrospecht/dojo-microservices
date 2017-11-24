package router

import (
	"net/http"

	"dojo-microservices/commons"

	"dojo-microservices/participantes/routes"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes.ParticipantesRoutes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = commons.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
