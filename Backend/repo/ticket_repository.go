package repo

import (
	"context"
	"flightbooking-app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TicketRepository struct {
	Collection *mongo.Collection
}

func connectToBase() {

}

func (repo *TicketRepository) GetTIcketsByUser(id primitive.ObjectID) ([]*model.Ticket, error) {

	filter := bson.M{"user._id": id}

	cursor, findErr := repo.Collection.Find(context.Background(), filter)
	if findErr != nil {
		return nil, findErr
	}

	var tickets []*model.Ticket

	for cursor.Next(context.Background()) {
		var ticket model.Ticket
		if err := cursor.Decode(&ticket); err != nil {
			return nil, err
		}
		tickets = append(tickets, &ticket)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}
