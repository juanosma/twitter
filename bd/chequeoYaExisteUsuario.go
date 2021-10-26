package bd

import (
	"context"
	"time"

	"github.com/juanosma/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*Funcion que chequea usuario*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//Se ejecuta al final de la rutina
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")
	//debe ser en json por eso {}
	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	//trae un objeto se convierte a hexadecimal string
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
