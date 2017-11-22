package handler

import (
	"dojo-microservices/util/dto"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func All(w http.ResponseWriter, r *http.Request) {
	participantes := dto.Participantes{
		dto.Participante{Nome: "Suenaga", Cpf: "000.000.000-00"},
		dto.Participante{Nome: "TESTE", Cpf: "111.111.111-11"},
	}

	json.NewEncoder(w).Encode(participantes)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	participanteId := vars["participanteId"]

	fmt.Fprintln(w, "Id:", participanteId)
}
