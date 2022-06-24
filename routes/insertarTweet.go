package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/antorez101/go-web-curso/bd"
	"github.com/antorez101/go-web-curso/models"
)

func InsertarTweet(response http.ResponseWriter, request *http.Request) {

	var t models.TweetDTO
	err := json.NewDecoder(request.Body).Decode(&t)

	if err != nil {
		http.Error(response, "Error decoding body "+err.Error(), 400)
	}
	tweet := models.Tweet{
		UserId:  IDUsuario,
		Mensaje: t.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertarTweet(tweet)

	if err != nil {
		http.Error(response, "Error insertando tweet "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(response, "Error al insertar el tweet", 400)
		return
	}

	response.WriteHeader(http.StatusCreated)
}
