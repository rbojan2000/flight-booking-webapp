package service

import (
	"flightbooking-app/model"
	"flightbooking-app/repo"
)

type FlightService struct {
	FlightRepo *repo.FlightRepository
}

func (service *FlightService) GetFlightPrice(flight *model.Flight) error {
	err := service.FlightRepo.GetTicketPrice(flight)
	if err != nil {
		return err
	}
	return nil
}
