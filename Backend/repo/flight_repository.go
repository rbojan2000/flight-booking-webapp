package repo

import (
	"context"
	"flightbooking-app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type FlightRepository struct {
	Collection *mongo.Collection
}

func (repo *FlightRepository) Create(flight *model.Flight) error {
	_, err := repo.Collection.InsertOne(context.Background(), &flight)
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

func (repo *FlightRepository) Delete(id primitive.ObjectID) (int64, error) {
	filter := bson.M{"_id": id}
	result, err := repo.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
func (repo *FlightRepository) GetTicketPrice(flight *model.Flight) error {
	return nil
}
