package model

import (
	"time"
)

type Flight struct {
	ID             uint      `json:"flightID" gorm:"primaryKey;autoIncrement"`
	Date           time.Time `json:"date"`
	DepartureID    uint      `json:"departureID" `
	ArrivalID      uint      `json:"arrivalID" `
	Departure      Location  `json:"departure" gorm:"foreignKey:DepartureID"`
	Arrival        Location  `json:"arrival" gorm:"foreignKey:ArrivalID"`
	PassengerCount int       `json:"passengerCount"`
	Capacity       int       `json:"capacity"`
}
