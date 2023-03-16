package handler

import (
	"encoding/json"
	"flightbooking-app/model"
	"flightbooking-app/service"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
}

func (handler *UserHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserService.Create(&user)
	if err != nil {
		println(err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
