package routes

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antorez101/go-web-curso/bd"
	"github.com/antorez101/go-web-curso/models"
	"github.com/dgrijalva/jwt-go"
)

var Email string
var IDUsuario string

func ProcesoToken(jwtKey string) (*models.Claims, bool, string, error) {
	clave := []byte("clave_privada")
	claims := &models.Claims{}

	splitToken := strings.Split(jwtKey, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	jwtKey = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(jwtKey, claims, func(t *jwt.Token) (interface{}, error) {
		return clave, nil
	})
	if err == nil {
		_, encontrado, ID := bd.ExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
			fmt.Println(ID)
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}
	return claims, false, string(""), err
}
