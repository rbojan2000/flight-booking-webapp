package utils

import (
	"flightbooking-app/model"
	"flightbooking-app/repo"
	"time"
)

type DatabaseSeeds struct {
	UserRepo   *repo.UserRepository
	FlightRepo *repo.FlightRepository
}

func (d *DatabaseSeeds) SeedData() error {

	flight1 := model.Flight{
		Date:           time.Now(),
		Departure:      model.Location{Country: "Ireland", City: "Dublin"},
		Arrival:        model.Location{Country: "France", City: "Paris"},
		Capacity:       230,
		Price:          123,
		PassengerCount: 120,
	}

	flight2 := model.Flight{
		Date:           time.Now(),
		Departure:      model.Location{Country: "Amsterdam", City: "Netherlands"},
		Arrival:        model.Location{Country: "Zagreb", City: "Croatia"},
		Capacity:       160,
		Price:          228,
		PassengerCount: 37,
	}

	tickets := []model.Ticket{}
	tickets = append(tickets, model.Ticket{
		Flight: model.Flight{
			Date:      time.Now(),
			Departure: model.Location{Country: "Amsterdam", City: "Netherlands"},
			Arrival:   model.Location{Country: "Zagreb", City: "Croatia"},
			Price:     228,
		},
	})

	user1 := model.User{
		Name:    "Marko",
		Surname: "Nikolic",
		Email:   "marko1@mail.com",
		Tickets: tickets,
	}

	d.FlightRepo.Create(&flight1)
	d.FlightRepo.Create(&flight2)
	d.UserRepo.Create(&user1)

	return nil
}
