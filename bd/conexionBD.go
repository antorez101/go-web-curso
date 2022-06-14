package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://mhernandez:MarkoMongo70@gowebc1.sbva9.mongodb.net/?retryWrites=true&w=majority")

/*
	Conectar a base de datos mongodb
*/
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
	log.Println("Conexion exitosa a MongoDB")
	return client
}

func ValidarConexion() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
