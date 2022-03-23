package main

import (
	"log"

	"github.com/gabrielnocce/projectx/db"
	"github.com/gabrielnocce/projectx/handlers"
)

func main() {

	if db.TestConnection() == 0 {
		log.Fatal("No hay conexion a la base de datos")
		return
	}
	handlers.HandlersFx()

}
