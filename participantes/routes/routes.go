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

var ParticipantesRoutes = Routes{
	Route{
		"All",
		"GET",
		"/",
		handler.All,
	},
	Route{
		"Show",
		"GET",
		"/participante/{participanteId}",
		handler.Show,
	},
}
