package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Flight Flight             `bson:"flight,omitempty"`
	User   User               `bson:"user,omitempty"`
	Price  float64            `bson:"price,omitempty"`
}
