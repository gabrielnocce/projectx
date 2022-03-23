package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Usuario es el modelo de la base de datos Mongo DB
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"nombre,omitempty"`
	LastName    string             `bson:"lastname" json:"lastname,omitempty"`
	DatoOfBirth time.Time          `bson:"dateOfBirth" json:"dateofbirth,omitempty"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password,omitempty"`
	Avatar      string             `bson:"avatar" json:"avatar,omitempty"`
	Banner      string             `bson:"banner" json:"banner,omitempty"`
	Biography   string             `bson:"biography" json:"biography,omitempty"`
}
