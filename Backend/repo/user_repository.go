package repo

import (
	"context"
	"flightbooking-app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (repo *UserRepository) Create(user *model.User) error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	user.Password, _ = encryptPassword(user.Password)

    _, err := repo.Collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return err
	}

	_, error := repo.Collection.InsertOne(context.Background(), &user)
	if error != nil {
		if mongo.IsDuplicateKeyError(error) {
			return error
		}
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

func (repo *UserRepository) GetTicketsByUserID(userID primitive.ObjectID) ([]model.Ticket, error) {
	filter := bson.M{"_id": userID}
	result := repo.Collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var user model.User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	return user.Tickets, nil
}

func encryptPassword(password string) (string, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPasswordBytes), nil
}
