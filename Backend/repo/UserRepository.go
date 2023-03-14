package repo

import (
	"context"
	"flightbooking-app/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (repo *UserRepository) Create(user *model.User) error {
	_, err := repo.Collection.InsertOne(context.Background(), &user)
	if err != nil {
		return err
	}
	return nil
}
