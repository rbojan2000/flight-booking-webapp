package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Flight struct {
	ID             uuid.UUID `json:"flightID"`
	Date           time.Time `json:"date"`
	Departure      Location  `json:"departure" gorm:"foreignKey:locationID;references:ID"`
	Arrival        Location  `json:"arrival"`
	PassengerCount int       `json:"passengerCount"`
	Capacity       int       `json:"capacity"`
}

func (flight *Flight) BeforeCreate(scope *gorm.DB) error {
	flight.ID = uuid.New()
	return nil
}
