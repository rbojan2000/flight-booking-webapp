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

	flight, err := service.FlightRepo.GetById(flightID)

	if err != nil {
		return fmt.Errorf(fmt.Sprintf("There is no flight on relation %s - %s.", flight.Arrival.City, flight.Departure.City))
	}

	if flight.PassengerCount+int64(numberOfTickets) > flight.Capacity {
		return fmt.Errorf(fmt.Sprintf("You can not buy %d tickets! There is only left %d tikets!", numberOfTickets, flight.Capacity-flight.PassengerCount))

	}

	_, err = service.FlightRepo.Collection.UpdateByID(context.Background(), flightID, bson.M{"$set": bson.M{"passengerCount": flight.PassengerCount + int64(numberOfTickets)}})
	if err != nil {
		return fmt.Errorf("error updating flight passenger count: %s", err.Error())
	}

	ticket := model.Ticket{
		Flight: *flight,
	}

	for i := 0; i < numberOfTickets; i++ {
		service.UserRepo.AddTicketToUser(userID, ticket)
	}

	return err
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
