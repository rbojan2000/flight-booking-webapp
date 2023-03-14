package main

import (
	"context"
	"flightbooking-app/handler"
	"flightbooking-app/repo"
	"flightbooking-app/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func initDB() *mongo.Client {
	database, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		print(err)
		return nil
	}

	return database
}

func startServer(handler *handler.UserHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/registerUser", handler.Create).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	client := initDB()
	if client == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	userRepo := &repo.UserRepository{Collection: client.Database("xws").Collection("users")}
	userService := &service.UserService{UserRepo: userRepo}
	userHandler := &handler.UserHandler{UserService: userService}


	startServer(userHandler)
}
