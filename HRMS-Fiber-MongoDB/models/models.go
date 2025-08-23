package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Salary float64            `json:"salary" bson:"salary"`
	Age    int                `json:"age" bson:"age"`
}
