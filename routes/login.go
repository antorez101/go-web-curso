package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/antorez101/go-web-curso/bd"
	"github.com/antorez101/go-web-curso/jwt"
	"github.com/antorez101/go-web-curso/models"
)

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		http.Error(response, "Usuario y/o contrase;a invalidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(response, "El email es requerido", 400)
		return
	}
	us, login := bd.IntentoLogin(t.Email, t.Password)
	if !login {
		http.Error(response, "Email y/o contrase;a invalidos", 400)
		return
	}
	jwtKey, err := jwt.GenerarToken(us)
	if err != nil {
		http.Error(response, "Error generando token "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusAccepted)
	json.NewEncoder(response).Encode(resp)

	// setting cookies
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(response, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
