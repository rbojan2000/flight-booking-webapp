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

	router.HandleFunc("/registerUser", handler.Create).Methods("POST")
	router.HandleFunc("/loginUser", handler.Login).Methods("POST")
	router.HandleFunc("/buyTicket", handler.BuyTicket).Methods("POST")

	router.HandleFunc("/userTickets/{id}", (handler.GetUserTickets)).Methods("GET")

	router.HandleFunc("/flights", flightHandler.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/flights/getAll", flightHandler.GetAll).Methods("GET")
	router.HandleFunc("/flights/{id}", flightHandler.GetById).Methods("GET")
	router.HandleFunc("/flights/{id}", flightHandler.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
