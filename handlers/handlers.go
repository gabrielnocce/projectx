package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gabrielnocce/projectx/middlew"
	"github.com/gabrielnocce/projectx/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Funcion que escucha el puerto y redirecciona el trafico
func HandlersFx() {

	router := mux.NewRouter()
	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.CheckDB(middlew.ValidateJWT(routers.ViewProfile))).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {

		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
