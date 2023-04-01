package utils

import (
	"context"
	"flightbooking-app/model"
	"flightbooking-app/repo"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type DatabaseSeeds struct {
	UserRepo   *repo.UserRepository
	FlightRepo *repo.FlightRepository
}

func (d *DatabaseSeeds) SeedData() error {

	flight1 := model.Flight{
		Date:           time.Date(2024, time.September, 31, 15, 45, 0, 0, time.UTC),
		Departure:      model.Location{Country: "Ireland", City: "Dublin"},
		Arrival:        model.Location{Country: "France", City: "Paris"},
		Capacity:       230,
		Price:          123,
		PassengerCount: 120,
		Available:      true,
	}

	flight2 := model.Flight{
		Date:           time.Date(2023, time.July, 12, 11, 33, 0, 0, time.UTC),
		Departure:      model.Location{Country: "Netherlands", City: "Amsterdam"},
		Arrival:        model.Location{Country: "Croatia", City: "Zagreb"},
		Capacity:       160,
		Price:          228,
		PassengerCount: 37,
		Available:      true,
	}

	flight3 := model.Flight{
		Date:           time.Date(2022, time.July, 12, 11, 33, 0, 0, time.UTC),
		Departure:      model.Location{Country: "New York", City: "USA"},
		Arrival:        model.Location{Country: "Poland", City: "Warsaw"},
		Capacity:       160,
		Price:          228,
		PassengerCount: 37,
		Available:      false,
	}

	tickets := []model.Ticket{}
	tickets = append(tickets, model.Ticket{
		Flight: model.Flight{
			Date:           time.Now(),
			Departure:      model.Location{Country: "Amsterdam", City: "Netherlands"},
			Arrival:        model.Location{Country: "Zagreb", City: "Croatia"},
			Price:          228,
			Capacity:       160,
			PassengerCount: 37,
			Available:      true,
		},
	})

	user1 := model.User{
		Name:    "Marko",
		Surname: "Nikolic",
		Email:   "marko1@mail.com",
		Tickets: tickets,
	}

	count, _ := d.FlightRepo.Collection.CountDocuments(context.Background(), bson.M{})

	if count == 0 {
		d.FlightRepo.Create(&flight1)
		d.FlightRepo.Create(&flight2)
		d.FlightRepo.Create(&flight3)
	}

	d.UserRepo.Create(&user1)

	return nil
}
