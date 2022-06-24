package bd

import (
	"context"
	"time"

	"github.com/antorez101/go-web-curso/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertarTweet(tweet models.Tweet) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := MongoC.Database("go-web-db")
	col := db.Collection("tweet")

	registro := bson.M{
		"userId":  tweet.UserId,
		"mensaje": tweet.Mensaje,
		"fecha":   tweet.Fecha,
	}
	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
