package main

import (
	"context"
	"flightbooking-app/handler"
	"flightbooking-app/repo"
	"flightbooking-app/server"
	"flightbooking-app/service"
	"flightbooking-app/utils"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	utils.LoadEnv()
	client := server.InitDB()
	if client == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	flightRepo := &repo.FlightRepository{Collection: client.Database("xws").Collection("flights")}
	flightService := &service.FlightService{FlightRepo: flightRepo}
	flightHandler := &handler.FlightHandler{FlightService: flightService}
	userRepo := &repo.UserRepository{Collection: client.Database("xws").Collection("users")}
	userService := &service.UserService{UserRepo: userRepo, FlightRepo: flightRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	data := &utils.DatabaseSeeds{FlightRepo: flightRepo, UserRepo: userRepo}
	data.SeedData()

	server.StartServer(userHandler, flightHandler)
}
