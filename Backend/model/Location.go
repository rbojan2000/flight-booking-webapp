package model

type Location struct {
	ID      uint   `json:"locationID" gorm:"primaryKey;autoIncrement"`
	Country string `json:"country" gorm:"not null;type:string"`
	City    string `json:"city" gorm:"not null;type:string"`
}
