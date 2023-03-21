package service

import (
	"context"
	"flightbooking-app/model"
	"flightbooking-app/repo"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	UserRepo   *repo.UserRepository
	FlightRepo *repo.FlightRepository
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

func (service *UserService) AssignTicketToUser(userID primitive.ObjectID, flightID primitive.ObjectID, numberOfTickets int) error {

	return nil
}

func (service *UserService) Login(user *model.User) (*model.User, error) {
	filter := bson.M{"email": user.Email}

	var result model.User
	err := service.UserRepo.Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	return &result, nil
}
