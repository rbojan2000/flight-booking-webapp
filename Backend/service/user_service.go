package service

import (
	"flightbooking-app/model"
	"flightbooking-app/repo"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	UserRepo *repo.UserRepository
}

func (service *UserService) Create(user *model.User) error {
	err := service.UserRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) GetTicketsByUserID(id primitive.ObjectID) ([]model.Ticket, error) {
	tickets, err := service.UserRepo.GetTicketsByUserID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("No tickets for user"))
	}
	return tickets, nil
}
