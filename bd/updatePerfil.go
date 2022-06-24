package bd

import (
	"context"
	"time"

	"github.com/antorez101/go-web-curso/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdatePerfil(perfil models.Usuario, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	bd := MongoC.Database("go-web-db")
	col := bd.Collection("usuarios")

	registro := make(map[string]interface{})

	if len(perfil.Nombre) > 0 {
		registro["nombre"] = perfil.Nombre
	}

	if len(perfil.Apellidos) > 0 {
		registro["apellidos"] = perfil.Apellidos
	}

	if len(perfil.Avatar) > 0 {
		registro["avatar"] = perfil.Avatar
	}

	if len(perfil.Banner) > 0 {
		registro["banner"] = perfil.Banner
	}

	if len(perfil.Biografia) > 0 {
		registro["biografia"] = perfil.Biografia
	}

	if len(perfil.Email) > 0 {
		registro["email"] = perfil.Email
	}

	if len(perfil.SitioWeb) > 0 {
		registro["sitioWeb"] = perfil.SitioWeb
	}

	if len(perfil.Ubicacion) > 0 {
		registro["ubicacion"] = perfil.Ubicacion
	}

	registro["fechaNacimiento"] = perfil.FechaNacimiento

	updatedString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, condicion, updatedString)

	if err != nil {
		return false, err
	}
	return true, nil
}
