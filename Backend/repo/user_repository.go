package repo

import (
	"context"
	"flightbooking-app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repo *UserRepository) FindByID(id primitive.ObjectID) (*model.User, error) {
	filter := bson.M{"_id": id}
	result := repo.Collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var user model.User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindAll() ([]*model.User, error) {
	cursor, err := repo.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*model.User
	for cursor.Next(context.Background()) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
