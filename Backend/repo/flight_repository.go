package repo

import (
	"flightbooking-app/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type FlightRepository struct {
	Collection *mongo.Collection
}

func (repo *FlightRepository) GetTicketPrice(flight *model.Flight) error {
	return nil
}
