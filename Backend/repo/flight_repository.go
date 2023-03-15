package repo

import (
	"flightbooking-app/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type FlightRepository struct {
	Collection *mongo.Collection
}

func (repo *FlightRepository) GetTicketPrice(flight *model.Flight) error {
	// filter := bson.M{
	// 	"departure.city": flight.Departure.City,
	// 	"arrival.city":   flight.Arrival.City,
	// 	"date":           flight.Date,
	// }

	// Dohvati let iz baze
	// var result model.Flight
	// if err := repo.collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
	// 	return 0, err
	// }

	return nil
}
