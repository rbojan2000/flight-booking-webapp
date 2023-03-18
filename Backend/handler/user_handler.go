package handler

import (
	"encoding/json"
	"flightbooking-app/model"
	"flightbooking-app/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	UserService *service.UserService
}

func (handler *UserHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserService.Create(&user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			writer.WriteHeader(http.StatusExpectationFailed)
			writer.Write([]byte("Email already exists"))
			return
		}
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) Login(writer http.ResponseWriter, req *http.Request) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	usr, err := handler.UserService.Login(&user)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": usr.ID.Hex(),
		"exp":    time.Now().Add(time.Minute * 30).Unix(),
		"userType": usr.Type,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Authorization", "Bearer " + tokenString)
	writer.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) GetUserTickets(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	log.Printf("Flight with id %s", id)

	tickets, err := handler.UserService.GetTicketsByUserID(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tickets)

}
