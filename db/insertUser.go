package db

import (
	"context"
	"time"

	"github.com/gabrielnocce/projectx/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Registro usuario en la base de datos. Ultima parada
func InsertUsr(u models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	db := MongoCN.Database("projectx")
	col := db.Collection("users")
	u.Password, _ = EncryptPass(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
