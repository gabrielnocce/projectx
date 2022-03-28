package db

import "golang.org/x/crypto/bcrypt"

//Funcion que encripta password
func EncryptPass(pass string) (string, error) {

	costo := 8 // la cantidad de encriptaciones q hace sobre la key. recomendado super admin 8 usuario normal 6
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
