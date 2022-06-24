package routes

import (
	"encoding/json"
	"net/http"

	"github.com/antorez101/go-web-curso/bd"
	"github.com/antorez101/go-web-curso/models"
)

func UpdatingPerfil(response http.ResponseWriter, request *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		http.Error(response, "Error decoding perfil "+err.Error(), 400)
		return
	}

	status, err := bd.UpdatePerfil(t, IDUsuario)

	if err != nil {
		http.Error(response, "Error updating perfil "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(response, "Error updating perfil ", 400)
		return
	}
	response.WriteHeader(http.StatusCreated)

}
