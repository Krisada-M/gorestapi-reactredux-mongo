package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Incomemodel struct {
	ID          primitive.ObjectID `bson:"_id"`
	Category    string             `json:"category"`
	Person      *string            `json:"person"`
	Productname *string            `json:"productname"`
	Purchase    *int               `json:"purchase"`
	Day         string             `json:"day"`
	Month       string             `json:"month"`
	Year        string             `json:"year"`
	Time        string             `json:"time"`
}
