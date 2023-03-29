package repo

import (
	"context"
	"errors"
	"flightbooking-app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type FlightRepository struct {
	Collection *mongo.Collection
}

func (repo *FlightRepository) Create(flight *model.Flight) error {
	filter := bson.M{"departure": flight.Departure, "arrival": flight.Arrival, "date": flight.Date}
	count, err := repo.Collection.CountDocuments(context.Background(), &filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("Flight with the same Departure, Arrival and Date already exists")
	}

	_, err = repo.Collection.InsertOne(context.Background(), &flight)
	if err != nil {
		return err
	}
	return nil
}

func (repo *FlightRepository) GetById(flightId primitive.ObjectID) (*model.Flight, error) {
	var f model.Flight
	filter := bson.M{"_id": flightId}
	err := repo.Collection.FindOne(context.Background(), filter).Decode(&f)
	if err != nil {
		return nil, err
	}
	return &f, err
}

func (repo *FlightRepository) SoftDelete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"available": false}}
	_, err := repo.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (repo *FlightRepository) GetFlightByArrivalDeppartureAndDate(arrivalCity string, departureCity string, date time.Time) (*model.Flight, error) {

	filter := bson.M{
		"departure.city": departureCity,
		"arrival.city":   arrivalCity,
		"date": bson.M{
			"$gte": date,
			"$lt":  date.AddDate(0, 0, 1),
		},
	}

	var flight model.Flight

	err := repo.Collection.FindOne(context.Background(), filter).Decode(&flight)

	return &flight, err
}

func (repo *FlightRepository) FindAll() ([]*model.Flight, error) {
	cursor, err := repo.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var flights []*model.Flight
	for cursor.Next(context.Background()) {
		var flight model.Flight
		if err := cursor.Decode(&flight); err != nil {
			return nil, err
		}
		flights = append(flights, &flight)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return flights, nil
}

func (repo *FlightRepository) FindAllAvailable() ([]*model.Flight, error) {
	filter := bson.M{"available": bson.M{"$ne": false}}
	cursor, err := repo.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var flights []*model.Flight
	for cursor.Next(context.Background()) {
		var flight model.Flight
		if err := cursor.Decode(&flight); err != nil {
			return nil, err
		}
		flights = append(flights, &flight)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return flights, nil
}
