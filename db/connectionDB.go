package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConnectDB()
var clientOtions = options.Client().ApplyURI("mongodb+srv://gnocce:Somepass2!@cluster0.q5zad.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

//Funcion que permite conectar con la base de datos
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOtions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Se establecio la conexion con la base de datos")
	return client
}

//Funcion de testo de conexion.
func TestConnection() int {

	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {

		return 0
	}
	return 1
}
