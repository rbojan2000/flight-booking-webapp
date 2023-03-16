package service

import (
	"flightbooking-app/model"
	"flightbooking-app/repo"
	"flightbooking-app/utils"
	"fmt"
	"time"

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

func (service *FlightService) GetTicketPrice(arrivalCity string, departureCity string, date string, ticketNum int) (float64, error) {
	dateFormatter := utils.DateFormatter{Format: time.RFC3339}
	parsedDate, err := dateFormatter.ParseYearMonthDayOfDateString(date)

	if err != nil {
		return -1, fmt.Errorf(fmt.Sprintf("Wrong format date!"))
	}

	flight, err := service.FlightRepo.GetFlightByArrivalDeppartureAndDate(arrivalCity, departureCity, parsedDate)

	if err != nil {
		return -1, fmt.Errorf(fmt.Sprintf("There is no flight on relation %s - %s.", arrivalCity, departureCity))
	}

	if flight.PassengerCount+ticketNum > flight.Capacity {
		return -1, fmt.Errorf(fmt.Sprintf("You can not buy %d tickets! There is only left %d tikets!", ticketNum, flight.Capacity-flight.PassengerCount))
	}

	price := float64(ticketNum) * flight.Price

	return price, err
}

func (service *FlightService) GetAll() ([]*model.Flight, error) {
	flights, err := service.FlightRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("There are no flights!"))
	}
	return flights, nil
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
