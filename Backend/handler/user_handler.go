package handler

import (
	"encoding/json"
	"flightbooking-app/model"
	"flightbooking-app/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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
	user.Type = 1
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
		"userID":   usr.ID.Hex(),
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
		"email": usr.Email,
		"role": usr.Type,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Authorization", "Bearer "+tokenString)
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(tokenString))
}

func (handler *UserHandler) BuyTicket(writer http.ResponseWriter, req *http.Request) {
	parseErr := req.ParseMultipartForm(32 << 20) // parse request body with max memory of 32 MB

	if parseErr != nil {
		http.Error(writer, parseErr.Error(), http.StatusBadRequest)
		return
	}

	userID := req.FormValue("userID")
	numberOfTickets := req.FormValue("numberOfTickets")
	flightID := req.FormValue("flightID")

	objectUserID, _ := primitive.ObjectIDFromHex(userID)
	objectFlightID, _ := primitive.ObjectIDFromHex(flightID)

	numOfTickets, _ := strconv.Atoi(numberOfTickets)

	err := handler.UserService.AssignTicketToUser(objectUserID, objectFlightID, numOfTickets)

	if err != nil {
		fmt.Print(err.Error())
		writer.WriteHeader(http.StatusConflict)
		writer.Write([]byte(err.Error())) // Write the error message to the response body

		return
	}

	writer.WriteHeader(http.StatusCreated)
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
