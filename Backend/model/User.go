package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty" `
	Surname  string             `bson:"surname,omitempty" `
	Email    string             `bson:"email,omitempty" `
	Password string             `bson:"password,omitempty"`
	Type     UserType           `bson:"type,omitempty"`
	Address  Location           `bson:"address,omitempty"`
}
