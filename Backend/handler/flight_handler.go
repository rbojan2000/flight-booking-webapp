package handler

import (
	"encoding/json"
	"flightbooking-app/model/dto"
	"flightbooking-app/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightHandler struct {
	FlightService *service.FlightService
}

func (handler *FlightHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	flights, err := handler.FlightService.GetAll()

	if err != nil {
		println("Error while getting flights")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(flights)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *FlightHandler) GetFlightPrice(writer http.ResponseWriter, req *http.Request) {
	var ticketPriceDTO dto.TicketPriceDTO

	err := json.NewDecoder(req.Body).Decode(&ticketPriceDTO)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	arrivalCity := ticketPriceDTO.ArrivalCity
	departureCity := ticketPriceDTO.DepartureCity
	date := ticketPriceDTO.Date
	ticketNum := ticketPriceDTO.TicketNum

	price, err := handler.FlightService.GetTicketPrice(arrivalCity, departureCity, date, ticketNum)
	fmt.Println(price)
	fmt.Println(err)

	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(price)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

func (handler *FlightHandler) Create(writer http.ResponseWriter, req *http.Request) {
	enableCors(&writer)
	var flightDto dto.FlightDTO
	err := json.NewDecoder(req.Body).Decode(&flightDto)

	err = handler.FlightService.Create(&flightDto)

	if err != nil {
		fmt.Println("Error:", err)
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *FlightHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	log.Printf("Flight with id %s", id)
	flight, err := handler.FlightService.GetById(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(flight)
}

func (handler *FlightHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	deletedCount, err := handler.FlightService.Delete(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if deletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
