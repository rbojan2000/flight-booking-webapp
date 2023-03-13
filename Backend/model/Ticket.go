package model

type Ticket struct {
	ID       uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	FlightId uint    `json:"flightId"`
	UserId   uint    `json:"userId"`
	Price    float64 `json:"price"`
}
