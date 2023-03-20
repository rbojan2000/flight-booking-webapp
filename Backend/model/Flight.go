package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Date           time.Time          `bson:"date,omitempty"`
	Departure      Location           `bson:"departure,omitempty"`
	Arrival        Location           `bson:"arrival,omitempty"`
<<<<<<< Updated upstream
	PassengerCount int                `bson:"passengerCount,omitempty"`
	Capacity       int                `bson:"capacity,omitempty"`
=======
	PassengerCount int64              `bson:"passengerCount,omitempty"`
	Capacity       int64              `bson:"capacity,omitempty"`
>>>>>>> Stashed changes
	Price          float64            `bson:"price,omitempty"`
}
