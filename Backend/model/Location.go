package model

type Location struct {
	Country string `bson:"country,omitempty"`
	City    string `bson:"city,omitempty"`
}
