package repo

import (
	"flightbooking-app/model"

	"gorm.io/gorm"
)

type FlightRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *FlightRepository) CreateFlight(flight *model.Flight) error {
	dResult := *repo.DatabaseConnection.Create(&flight.Departure)
	if dResult.Error != nil {
		return dResult.Error
	}

	aResult := *repo.DatabaseConnection.Create(&flight.Arrival)
	if aResult.Error != nil {
		return aResult.Error
	}
	//flight.DepartureID = flight.Departure.ID
	//flight.ArrivalID = flight.Arrival.ID

	dbResult := *repo.DatabaseConnection.Create(flight)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *FlightRepository) FindById(id string) (model.Flight, error) {
	flight := model.Flight{}
	dbResult := repo.DatabaseConnection.First(&flight, "id = ?", id)
	if dbResult != nil {
		return flight, dbResult.Error
	}
	return flight, nil
}
