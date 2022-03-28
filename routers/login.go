package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gabrielnocce/projectx/db"
	"github.com/gabrielnocce/projectx/jwt"
	"github.com/gabrielnocce/projectx/models"
)

// Router correspondiente al login

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		http.Error(w, "Usuario y/o contraseña incorrecta"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	document, exist := db.TryLogin(user.Email, user.Password)

	if exist == false {

		http.Error(w, "Usuario y/o contraseña incorrecta", 400)
		return
	}

	jwtkey, err := jwt.GenerateJWT(document)

	if err != nil {

		http.Error(w, "Ocurrio un error al intentar generar el token de seguridad"+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtkey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//Generar Cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtkey,
		Expires: expirationTime,
	})
}
