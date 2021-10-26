package bd

import (
	"context"
	"time"

	"github.com/juanosma/twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Inserto datos del usuario en BD*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//Se ejecuta al final de la rutina
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")
	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}
	ObjectID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjectID.String(), true, nil

}
