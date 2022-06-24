package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	UserId  string             `bson:"userId" json:"userId,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
