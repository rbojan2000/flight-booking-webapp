package handler

import (
	"encoding/json"
	"flightbooking-app/model"
	"flightbooking-app/service"
	"net/http"
)

type FlightHandler struct {
	FlightService *service.FlightService
}

func (handler *FlightHandler) GetFlightPrice(writer http.ResponseWriter, req *http.Request) {
	var flight model.Flight
	err := json.NewDecoder(req.Body).Decode(&flight)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		println("Error while encoding json")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = handler.FlightService.GetFlightPrice(&flight)
	if err != nil {
		println(err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
