package main

import (
	"dojo-microservices/commons"
	"dojo-microservices/participantes/handler"
	"dojo-microservices/participantes/router"
	"log"
	"net/http"
	"os"

	"github.com/creamdog/gonfig"
)

func main() {
	f, _ := os.Open("participantes/properties.yml")
	defer f.Close()
	conf, _ := gonfig.FromYml(f)
	handler.Config = commons.Config{Yml: conf}

	log.Fatal(http.ListenAndServe(":8080", router.NewRouter()))
}
