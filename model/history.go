package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type History struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Date  time.Time `json:"date" bson:"date"`
}
