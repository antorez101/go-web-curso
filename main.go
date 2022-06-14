package main

import (
	"log"

	"github.com/antorez101/go-web-curso/bd"
	"github.com/antorez101/go-web-curso/handlers"
)

func main() {
	if bd.ValidarConexion() == 0 {
		log.Fatal("Sin conexcion a BD")
		return
	}
	handlers.Manejadores()
}
