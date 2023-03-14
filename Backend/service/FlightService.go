package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type FlightService struct {
	FlightRepo *repo.FlightRepository
}

func (service *FlightService) Create(flight *model.Flight) error {
	err := service.FlightRepo.CreateFlight(flight)
	if err != nil {
		return err
	}
	return nil
}

func (service *FlightService) FindUser(id string) (*model.Flight, error) {
	flight, err := service.FlightRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Flight with id %s not found", id))
	}
	return &flight, nil
}
