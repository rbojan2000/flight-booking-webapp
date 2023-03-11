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

	database.AutoMigrate(&model.User{})

	database.Exec("INSERT IGNORE INTO users VALUES ('aec7e123-233d-4a09-a289-75308ea5b7e6', 'David')")
	return database
}

func startServer(handler *handler.UserHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/users", handler.Create).Methods("POST")

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
	repo := &repo.UserRepository{DatabaseConnection: database}
	service := &service.UserService{UserRepo: repo}
	handler := &handler.UserHandler{UserService: service}

	startServer(handler)
}
