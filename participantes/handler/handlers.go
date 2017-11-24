package handler

import (
	"dojo-microservices/commons/dto"
	"encoding/json"
	"net/http"

	"dojo-microservices/commons"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// TODO export to yml
var databaseUrl = "localhost"
var database = "participantes"
var collectionName = "participante"

func All(w http.ResponseWriter, r *http.Request) {
	ds := commons.GetDataStore(databaseUrl)
	defer ds.Session.Close()

	var participantes dto.Participantes

	// TODO HELPER ?
	collection := ds.GetCollection(database, collectionName)
	collection.Find(bson.M{}).All(&participantes)

	json.NewEncoder(w).Encode(participantes)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	participanteId := vars["participanteId"]

	ds := commons.GetDataStore(databaseUrl)
	defer ds.Session.Close()

	var participante dto.Participante
	collection := ds.GetCollection(database, collectionName)
	collection.FindId(bson.ObjectIdHex(participanteId)).One(&participante)

	json.NewEncoder(w).Encode(participante)
}
