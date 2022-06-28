package routes

import (
	"net/http"

	"github.com/antorez101/go-web-curso/bd"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El id no debe ser null", http.StatusBadRequest)
	}
	idUsuario := r.URL.Query().Get("idUsuario")
	if len(idUsuario) < 1 {
		http.Error(w, "El idUsuario no debe ser null ", http.StatusBadRequest)
	}

	err := bd.DeleteTweet(ID, idUsuario)
	if err != nil {
		http.Error(w, "Error deleting the tweet "+err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
