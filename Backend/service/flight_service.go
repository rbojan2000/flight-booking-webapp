package service

import (
	"flightbooking-app/model"
	"flightbooking-app/model/dto"
	"flightbooking-app/repo"
	"flightbooking-app/utils"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightService struct {
	FlightRepo *repo.FlightRepository
}

func (service *FlightService) Create(flightDto *dto.FlightDTO) error {
	var flight model.Flight
	flight.Arrival.City = flightDto.ArrivalCity
	flight.Arrival.Country = flightDto.ArrivalCountry
	flight.Departure.City = flightDto.DepartureCity
	flight.Departure.Country = flightDto.DepartureCountry
	flight.Price, _ = strconv.ParseFloat(flightDto.Price, 64)
	flight.PassengerCount, _ = strconv.ParseInt(flightDto.TicketNum, 10, 32)
	flight.Capacity, _ = strconv.ParseInt(flightDto.TicketNum, 10, 0)
	flight.Available = true
	date, _ := time.Parse("2006-01-02, 15:04", flightDto.DateAndTime)
	flight.Date = date

	if flight.Price <= 0 || flight.PassengerCount <= 0 {
		return fmt.Errorf("You can't enter negative number !")
	}

	err := service.FlightRepo.Create(&flight)
	if err != nil {
		return err
	}
	return nil
}

func (service *FlightService) GetTicketPrice(arrivalCity string, departureCity string, date string, ticketNum int64) (float64, error) {
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

func (service *FlightService) GetAllAvailable() ([]*model.Flight, error) {
	flights, err := service.FlightRepo.FindAllAvailable()
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

func (service *FlightService) Delete(id primitive.ObjectID) error {
	err := service.FlightRepo.SoftDelete(id)
	if err != nil {
		return err
	}
	return err
}
