package models

type TweetDTO struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
