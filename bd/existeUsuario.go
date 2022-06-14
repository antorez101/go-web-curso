package bd

import (
	"context"
	"time"

	"github.com/antorez101/go-web-curso/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ExisteUsuario(email string) (models.Usuario, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoC.Database("go-web-db")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	u := col.FindOne(ctx, condicion).Decode(&resultado)

	ID := resultado.ID.Hex()

	if u != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
