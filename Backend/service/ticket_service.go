package service

import (
	"flightbooking-app/model"
	"flightbooking-app/repo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketService struct {
	TicketRepo *repo.TicketRepository
}

func (service *TicketService) GetTicketsForUser(id primitive.ObjectID) ([]*model.Ticket, error) {
	tickets, err := service.TicketRepo.GetTIcketsByUser(id)
	if err != nil {
		return tickets, err
	}

	return tickets, err
}
