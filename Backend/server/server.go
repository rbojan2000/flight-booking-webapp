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
	router.HandleFunc("/userTickets/{id}", middleware.RequireAuth("1", handler.GetUserTickets)).Methods("GET")

	router.HandleFunc("/flights", middleware.RequireAuth("1", flightHandler.Create)).Methods("POST")
	router.HandleFunc("/flights/getAll", middleware.RequireAuth("1", flightHandler.GetAll)).Methods("GET")
	router.HandleFunc("/flights/{id}", middleware.RequireAuth("1", flightHandler.GetById)).Methods("GET")
	router.HandleFunc("/flights/{id}", middleware.RequireAuth("1", flightHandler.Delete)).Methods("DELETE")
	router.HandleFunc("/flights/getFlightPrice", middleware.RequireAuth("1", flightHandler.GetFlightPrice)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}