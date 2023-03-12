package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Flight struct {
	ID   uuid.UUID `json:"flightID"`
	Date time.Time `json:"date"`
	// LocationID1 string    `json:"-"`
	// LocationID2 string    `json:"-"`

	Departure      Location `json:"departure" gorm:"not null;type:string"`
	Arrival        Location `json:"arrival" gorm:"not null;type:string"`
	PassengerCount int      `json:"passengerCount"`
	Capacity       int      `json:"capacity"`
}

func (flight *Flight) BeforeCreate(scope *gorm.DB) error {
	flight.ID = uuid.New()
	return nil
}
