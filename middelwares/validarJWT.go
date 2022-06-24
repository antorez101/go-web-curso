package middelwares

import (
	"net/http"

	"github.com/antorez101/go-web-curso/routes"
)

func ValidaJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error validando jwt "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}

}
