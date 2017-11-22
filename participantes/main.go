package main

import (
	"log"
	"net/http"

	"dojo-microservices/participantes/router"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", router.NewRouter()))
}
