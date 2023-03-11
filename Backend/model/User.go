package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name" gorm:"not null;type:string"`
	Surname  string    `json:"surname" gorm:"not null;type:string"`
	Email    string    `json:"email" gorm:"not null;type:string"`
	Password string    `json:"password" gorm:"not null;type:string"`
	Type     UserType  `json:"type" gorm:"not null;type:int"`
	Address  Location  `json:"address" gorm:"not null;type:string"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.ID = uuid.New()
	return nil
}
