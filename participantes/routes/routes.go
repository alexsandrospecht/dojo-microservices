package routes

import (
	"net/http"
	"dojo-microservices/participantes/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var RoutesConstants = Routes{
	Route{
		"All",
		"GET",
		"/participantes",
		handler.All,
	},
	Route{
		"Show",
		"GET",
		"/participante/{participanteId}",
		handler.Show,
	},
}
