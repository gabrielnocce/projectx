package db

import (
	"github.com/gabrielnocce/projectx/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.User, bool) {

	usu, encontrado, _ := ExistUser(email)

	if encontrado == false {

		return usu, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {

		return usu, false
	}
	return usu, true

}
