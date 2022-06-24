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
	router.HandleFunc("/login", middelwares.CheckDataBase(routes.Login)).Methods("POST")
	router.HandleFunc("/verPerfil", middelwares.CheckDataBase(middelwares.ValidaJWT(routes.VerPerfil))).Methods("GET")
	router.HandleFunc("/updatePerfil", middelwares.CheckDataBase(middelwares.ValidaJWT(routes.UpdatingPerfil))).Methods("PUT")
	router.HandleFunc("/insertTweet", middelwares.CheckDataBase(middelwares.ValidaJWT(routes.InsertarTweet))).Methods("POST")
	router.HandleFunc("/getTweets", middelwares.CheckDataBase(middelwares.ValidaJWT(routes.GetTweets))).Methods("GET")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
