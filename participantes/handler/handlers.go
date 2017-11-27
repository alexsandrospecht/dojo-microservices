package handler

import (
	"dojo-microservices/commons"
	"dojo-microservices/commons/dto"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var Config commons.Config

func All(w http.ResponseWriter, r *http.Request) {
	ds := commons.GetDataStore(Config.GetValue("databaseUrl"))
	defer ds.Session.Close()

	var participantes dto.Participantes
	ds.FindAll(Config.GetValue("database"), Config.GetValue("collectionName"), bson.M{}, &participantes)

	json.NewEncoder(w).Encode(participantes)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	participanteId := vars["participanteId"]

	ds := commons.GetDataStore(Config.GetValue("databaseUrl"))
	defer ds.Session.Close()

	var participante dto.Participante
	ds.FindOne(Config.GetValue("database"), Config.GetValue("collectionName"), bson.ObjectIdHex(participanteId), &participante)

	json.NewEncoder(w).Encode(participante)
}
