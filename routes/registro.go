package routes

import (
	"encoding/json"
	"net/http"

	"github.com/antorez101/go-web-curso/bd"
	"github.com/antorez101/go-web-curso/models"
)

func Registro(response http.ResponseWriter, request *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		http.Error(response, "Error al decodificar objeto "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(response, "Email invalido ", 400)
		return
	}
	if len(t.Email) < 6 {
		http.Error(response, "Especificar contrase;na del almenos 6 caracteres ", 400)
		return
	}

	_, encontrado, _ := bd.ExisteUsuario(t.Email)
	if encontrado {
		http.Error(response, "El usuario ya existe con el email "+t.Email, 400)
		return
	}
	_, status, err := bd.InsertarUsuario(t)
	if err != nil {
		http.Error(response, "ERROR insertando usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(response, "No se inserto el usuario ", 400)
		return
	}
	response.WriteHeader(http.StatusCreated)

}
