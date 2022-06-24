package bd

import (
	"github.com/antorez101/go-web-curso/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, pass string) (models.Usuario, bool) {

	us, existe, _ := ExisteUsuario(email)
	if !existe {
		return us, false
	}
	password := []byte(pass)
	passbd := []byte(us.Password)
	err := bcrypt.CompareHashAndPassword(passbd, password)

	if err != nil {
		return us, false
	}
	return us, true
}
