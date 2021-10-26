package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://juanes125:Osma6720.@cluster0.rtj7h.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*Conectar BD es la funcion que me permite conectar a la bd*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Printf("Conexión exitosa con la BD")
	return client
}

/*Permite chequear la conexion*/
func ChequeoConnection() int {

	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
