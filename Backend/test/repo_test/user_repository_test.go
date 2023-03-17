package repo_test

import (
	"context"
	"flightbooking-app/model"
	"flightbooking-app/repo"
	"flightbooking-app/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var userRepoMock *repo.UserRepository

func setup() {
	utils.LoadEnv()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION_PORT")))
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

func TestSameEmailCreate(t *testing.T) {
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
	assert.Equal(t, true, mongo.IsDuplicateKeyError(userRepoMock.Create(user)))
}
