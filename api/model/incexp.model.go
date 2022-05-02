package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type IncExpmodel struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Category    string             `json:"category"`
	Person      *string            `json:"person"`
	Productname *string            `json:"productname"`
	Purchase    *int               `json:"purchase"`
	Date        Date               `json:"date"`
}

type Date struct {
	Day   string `json:"day"`
	Month string `json:"month"`
	Year  string `json:"year"`
	Time  string `json:"time"`
}
