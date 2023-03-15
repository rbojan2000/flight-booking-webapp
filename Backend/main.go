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

func startServer(handler *handler.UserHandler, flightHandler *handler.FlightHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/registerUser", handler.Create).Methods("POST")
	router.HandleFunc("/flights/getAll", flightHandler.GetAll).Methods("GET")
	router.HandleFunc("/flights/getFlightPrice", flightHandler.GetFlightPrice).Methods("POST")
	router.HandleFunc("/flights", flightHandler.Create).Methods("POST")
	router.HandleFunc("/flights/{id}", flightHandler.GetById).Methods("GET")
	router.HandleFunc("/flights/{id}", flightHandler.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	client := initDB()

	print("Server started")

	if client == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	userRepo := &repo.UserRepository{Collection: client.Database("xws").Collection("users")}
	flightRepo := &repo.FlightRepository{Collection: client.Database("xws").Collection("flights")}

	userService := &service.UserService{UserRepo: userRepo}
	flightService := &service.FlightService{FlightRepo: flightRepo}

	userHandler := &handler.UserHandler{UserService: userService}
	flightHandler := &handler.FlightHandler{FlightService: flightService}

	startServer(userHandler, flightHandler)
}
