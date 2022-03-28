package db

import (
	"context"
	"fmt"
	"time"

	"github.com/gabrielnocce/projectx/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Funcion que busca el prefil de un usuario en BD

func SearchProfile(ID string) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	db := MongoCN.Database("projectx")
	col := db.Collection("users")

	var usrProfile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&usrProfile)
	usrProfile.Password = ""
	if err != nil {

		fmt.Println("Registro no Encontrado" + err.Error())
		return usrProfile, err
	}

	return usrProfile, nil

}
