package repo_test

import (
	"context"
	"flightbooking-app/model"
	"flightbooking-app/repo"
	"flightbooking-app/utils"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var flightRepoMock *repo.FlightRepository

func setupFlightRepo() {
	utils.LoadEnv()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION_PORT")))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	flightRepoMock = &repo.FlightRepository{Collection: client.Database("test_db").Collection("flights")}
}

func TestCreateFlight(t *testing.T) {
	setupFlightRepo()
	defer flightRepoMock.Collection.Drop(context.Background())

	flight := &model.Flight{
		Date:           time.Date(2023, time.March, 12, 13, 0, 0, 0, time.UTC),
		Departure:      model.Location{Country: "Srbija", City: "Beograd"},
		Arrival:        model.Location{Country: "Italija", City: "Rim"},
		PassengerCount: 150,
		Capacity:       200,
		Price:          50,
	}

	flightRepoMock.Create(flight)
	flights, _ := flightRepoMock.FindAll()
	assert.Equal(t, flights[0].Date, flight.Date, "Two dates should be equal")

}
