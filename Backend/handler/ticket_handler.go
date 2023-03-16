package handler

import (
	"encoding/json"
	"flightbooking-app/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketHandler struct {
	TicketService *service.TicketService
}

func (handler *TicketHandler) GetTicketsForUser(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, err := primitive.ObjectIDFromHex(vars["id"])

	tickets, err := handler.TicketService.GetTicketsForUser(id)

	fmt.Println(tickets)

	if err != nil {
		println("Error while getting flights")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(tickets)
	writer.Header().Set("Content-Type", "application/json")
}
