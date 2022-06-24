package routes

import (
	"encoding/json"
	"net/http"

	"github.com/antorez101/go-web-curso/bd"
)

func VerPerfil(response http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(response, "Debe enviar el parametro ID ", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscarPerfil(ID)

	if err != nil {
		http.Error(response, "Error buscando usuario "+err.Error(), 400)
		return
	}
	response.Header().Set("context-type", "application/json")
	response.WriteHeader(http.StatusAccepted)
	json.NewEncoder(response).Encode(perfil)

}
