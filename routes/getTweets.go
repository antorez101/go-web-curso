package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/antorez101/go-web-curso/bd"
)

func GetTweets(response http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")

	if len(id) < 0 {
		http.Error(response, "Id no debe ser nulo ", http.StatusBadRequest)
		return
	}
	if len(request.URL.Query().Get("pagina")) < 0 {
		http.Error(response, "Pagina no debe ser nulo ", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(request.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(response, "Pagina debe ser mayor a 0 ", http.StatusBadRequest)
		return

	}
	pag := int64(pagina)
	resp, status := bd.GetTweets(id, pag)
	if !status {
		http.Error(response, "Error leyendo los tweets ", http.StatusBadRequest)
		return
	}
	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(resp)
}
