package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	ID       uuid.UUID `json:"id"`
	FlightId uuid.UUID `json:"flightId"`
	UserId   uuid.UUID `json:"userId"`
	Price    float64   `json:"price"`
}

func (ticket *Ticket) BeforeCreate(scope *gorm.DB) error {
	ticket.ID = uuid.New()
	return nil
}
