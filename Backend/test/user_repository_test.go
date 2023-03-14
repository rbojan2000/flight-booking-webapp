package user_repo_test

import (
	"context"
	"flightbooking-app/model"
	"flightbooking-app/repo"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var userRepoMock *repo.UserRepository

func setup() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	userRepoMock = &repo.UserRepository{Collection: client.Database("test_db").Collection("users")}
}

func TestCreateUser(t *testing.T) {
	setup()
	defer userRepoMock.Collection.Drop(context.Background())

	user := &model.User{
		Name:     "David",
		Surname:  "Mijailovic",
		Email:    "david@gmail.com",
		Password: "pass",
		Type:     1,
		Address: model.Location{
			Country: "Srbija",
			City:    "Novi Sad",
		},
	}

	userRepoMock.Create(user)
	users, _ := userRepoMock.FindAll()
	assert.Equal(t, users[0].Email, user.Email, "Two emails should be equal")
	
}