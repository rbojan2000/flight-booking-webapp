package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Country string             `bson:"country,omitempty"`
	City    string             `bson:"city,omitempty"`
}
