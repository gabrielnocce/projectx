package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"

	"time"

	"github.com/gabrielnocce/projectx/models"
)

func GenerateJWT(usu models.User) (string, error) {

	myKey := []byte("Somepass2!")

	payload := jwt.MapClaims{

		"email":       usu.Email,
		"name":        usu.Name,
		"lastName":    usu.LastName,
		"datoOfBirth": usu.DatoOfBirth,
		"biography":   usu.Biography,
		"_id":         usu.ID.Hex(),
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
