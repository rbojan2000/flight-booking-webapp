package server

import (
	"flightbooking-app/handler"
	"flightbooking-app/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer(handler *handler.UserHandler, flightHandler *handler.FlightHandler) {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.CorsMiddleware)

	router.HandleFunc("/registerUser", handler.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/loginUser", handler.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/buyTicket", handler.BuyTicket).Methods("POST")

	router.HandleFunc("/userTickets/{id}", (handler.GetUserTickets)).Methods("GET")

	router.HandleFunc("/flights", flightHandler.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/flights/getAll", flightHandler.GetAll).Methods("GET")
	router.HandleFunc("/flights/getAllAvailable", flightHandler.GetAllAvailable).Methods("GET")
	router.HandleFunc("/flights/getByParams/SearchDepartureCity={departure}&SearchArrivalCity={arrival}&SearchDate={date}&SearchPassengerCount={count}", flightHandler.GetByParams).Methods("GET")

	router.HandleFunc("/flights/getFreeFlights", flightHandler.GetFreeFlights).Methods("GET")

	router.HandleFunc("/flights/{id}", flightHandler.GetById).Methods("GET")
	router.HandleFunc("/flights/{id}", flightHandler.Delete).Methods("DELETE", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", router))
}