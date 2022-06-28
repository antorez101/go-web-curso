package bd

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(ID string, idPerfil string) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	bd := MongoC.Database("go-web-db")
	col := bd.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":    objID,
		"userid": idPerfil,
	}

	_, err := col.DeleteOne(ctx, condicion)
	log.Println("Deleting tweet")
	return err
}
