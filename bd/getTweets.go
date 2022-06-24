package bd

import (
	"context"
	"log"
	"time"

	"github.com/antorez101/go-web-curso/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(userId string, pagina int64) ([]*models.Tweets, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := MongoC.Database("go-web-db")
	col := db.Collection("tweet")

	var tweets []*models.Tweets

	condicion := bson.M{
		"userId": userId,
	}

	optiones := options.Find()
	optiones.SetLimit(20)
	optiones.SetSort(bson.D{{Key: "fecha", Value: -1}})

	cursor, err := col.Find(ctx, condicion, optiones)

	if err != nil {
		log.Fatalf("Error getting tweets %s ", err.Error())
		return tweets, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.Tweets

		err := cursor.Decode(&registro)
		if err != nil {
			return tweets, false
		}
		tweets = append(tweets, &registro)
	}
	return tweets, true
}
