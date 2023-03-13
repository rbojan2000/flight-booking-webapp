package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "root:test@tcp(localhost:3306)/flight_booking?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}
	database.AutoMigrate(&model.Location{})
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Flight{})
	database.AutoMigrate(&model.Ticket{})

	//database.Exec("INSERT IGNORE INTO users VALUES ('1', 'David', 'Mijailovic', 'david@mail.com', 'david', 3, '')")
	//database.Exec("INSERT IGNORE INTO flights VALUES ('aec7e123-233d-4a09-a289-75308ea5b7e6', '2023-03-13T15:30:00Z')")
	//database.Exec("INSERT IGNORE INTO locations VALUES ('aec7e123-233d-4a09-a289-75308ea5b7e6', 'David')")
	//database.Exec("INSERT IGNORE INTO tickets VALUES ('aec7e123-233d-4a09-a289-75308ea5b7e6', 'David')")

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
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	userRepo := &repo.UserRepository{DatabaseConnection: database}
	userService := &service.UserService{UserRepo: userRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	flightRepo := &repo.FlightRepository{DatabaseConnection: database}
	flightService := &service.FlightService{FlightRepo: flightRepo}
	flightHandler := &handler.FlightHandler{FlightService: flightService}

	startServer(userHandler, flightHandler)
}
