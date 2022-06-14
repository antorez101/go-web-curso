package bd

import (
	"context"
	"time"

	"github.com/antorez101/go-web-curso/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertarUsuario(u models.Usuario) (string, bool, error) {
	// Se adiciona al contexto un time de 15 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Se
	defer cancel()                                                           // Este cancel elimina del contexto el timeout creado de 15 segundos

	db := MongoC.Database("go-web-db")
	col := db.Collection("usuarios")

	u.Password, _ = EncryptPassword(u.Password)

	// el context "ctx" contiene el timeout de 15 segundos credo lineas arriba
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
