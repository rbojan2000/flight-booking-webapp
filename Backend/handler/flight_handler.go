package handler

import (
	"encoding/json"
	"flightbooking-app/model"
	"flightbooking-app/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FlightHandler struct {
	FlightService *service.FlightService
}

func (handler *FlightHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var flight model.Flight
	err := json.NewDecoder(req.Body).Decode(&flight)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	err = handler.FlightService.Create(&flight)
	if err != nil {
		println("Error while creating a new flight")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *FlightHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Flight with id %s", id)
	flight, err := handler.FlightService.FindUser(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(flight)
}
