package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gabrielnocce/projectx/db"
	"github.com/gabrielnocce/projectx/models"
)

var Email string
var IDUsuario string

//Procesa token para validar mensaje
func TokenProcess(tk string) (*models.Claims, bool, string, error) {
	myKey := []byte("Somepass2!")
	claims := &models.Claims{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de Token Invalido")

	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, encontrado, _ := db.ExistUser(claims.Email)

		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()

		}
		return claims, encontrado, IDUsuario, nil

	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token Invalido")
	}
	return claims, false, string(""), err

}
