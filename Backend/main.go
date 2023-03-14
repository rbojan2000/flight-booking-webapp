package main

import (
	"context"
	"database-example/handler"
	"database-example/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
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
	fmt.Println(database.Database("xws").Name())

	return database
}

func startServer(handler *handler.UserHandler, flightHandler *handler.FlightHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/users", handler.Create).Methods("POST")

	router.HandleFunc("/flights", flightHandler.Create).Methods("POST")
	router.HandleFunc("/flights/{id}", flightHandler.Get).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
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

	usersCollection := client.Database("xws").Collection("users")

	individual := model.User{
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

	_, err := usersCollection.InsertOne(context.TODO(), &individual)
	if err != nil {
		log.Fatalln("Error Inserting Document", err)
	}

	// retrieve single and multiple documents with a specified filter using FindOne() and Find()
	// create a search filer
	filter := bson.D{}

	// retrieve all the documents that match the filter
	cursor, err := usersCollection.Find(context.TODO(), filter)
	// check for errors in the finding
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var results []bson.D
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// display the documents retrieved
	fmt.Println("displaying all results from the search query")
	for _, result := range results {
		fmt.Println(result)
	}

	// userRepo := &repo.UserRepository{DatabaseConnection: database}
	// userService := &service.UserService{UserRepo: userRepo}
	// userHandler := &handler.UserHandler{UserService: userService}

	// flightRepo := &repo.FlightRepository{DatabaseConnection: database}
	// flightService := &service.FlightService{FlightRepo: flightRepo}
	// flightHandler := &handler.FlightHandler{FlightService: flightService}

	//startServer(userHandler, flightHandler)
}
