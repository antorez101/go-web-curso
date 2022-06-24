package jwt

import (
	"time"

	"github.com/antorez101/go-web-curso/models"
	"github.com/dgrijalva/jwt-go"
)

func GenerarToken(t models.Usuario) (string, error) {

	clave := []byte("clave_privada")

	claims := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellido":         t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioWeb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString(clave)
	if err != nil {
		return tokenSigned, err
	}
	return tokenSigned, nil
}
