package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/antorez101/go-web-curso/middelwares"
	"github.com/antorez101/go-web-curso/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {

	router := mux.NewRouter()

	router.HandleFunc("/registro", middelwares.CheckDataBase(routes.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
