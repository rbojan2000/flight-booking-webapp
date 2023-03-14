package service

import (
	"flightbooking-app/model"
	"flightbooking-app/repo"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightService struct {
	FlightRepo *repo.FlightRepository
}

func (service *FlightService) Create(flight *model.Flight) error {
	err := service.FlightRepo.Create(flight)
	if err != nil {
		return err
	}
	return nil
}

func (service *FlightService) GetById(id primitive.ObjectID) (*model.Flight, error) {
	flight, err := service.FlightRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Flight item with id %s not found", id))
	}
	return flight, nil
}

func (service *FlightService) Delete(id primitive.ObjectID) (int64, error) {
	deletedCount, err := service.FlightRepo.Delete(id)
	if err != nil {
		return 0, err
	}
	return deletedCount, err
}
