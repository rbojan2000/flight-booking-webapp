package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID      uuid.UUID `json:"locationID"`
	Country string    `json:"country" gorm:"not null;type:string"`
	City    string    `json:"city" gorm:"not null;type:string"`
}

func (location *Location) BeforeCreate(scope *gorm.DB) error {
	location.ID = uuid.New()
	return nil
}
