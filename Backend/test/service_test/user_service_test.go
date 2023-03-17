package user_service_test

import (
	"context"
	"flightbooking-app/model"
	"flightbooking-app/repo"
	"flightbooking-app/service"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var userServiceMock *service.UserService

func loadEnv() {
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func setup() {
	loadEnv()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION_PORT")))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	userRepoMock := &repo.UserRepository{Collection: client.Database("test_db").Collection("users")}
	userServiceMock = &service.UserService{UserRepo: userRepoMock}
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
}

func TestLoginUser(t *testing.T){
	setup()
	defer userServiceMock.UserRepo.Collection.Drop(context.Background())

	user := &model.User{
		Email:    "david@gmail.com",
		Password: "pass",
	}

	usr, err := userServiceMock.Login(user)
	assert.Equal(t, nil, err)
	assert.Equal(t, user.Email, usr.Email)
}

func TestLoginUserWrongPassword(t *testing.T){
	setup()
	defer userServiceMock.UserRepo.Collection.Drop(context.Background())

	user := &model.User{
		Email:    "david@gmail.com",
		Password: "gass",
	}

	_, err := userServiceMock.Login(user)
	assert.Equal(t, "crypto/bcrypt: hashedPassword is not the hash of the given password", err.Error())
}

func TestLoginUserWrongEmail(t *testing.T){
	setup()
	defer userServiceMock.UserRepo.Collection.Drop(context.Background())

	user := &model.User{
		Email:    "gas@gmail.com",
		Password: "gass",
	} 

	_, err := userServiceMock.Login(user)
	assert.Equal(t, "mongo: no documents in result", err.Error())
}