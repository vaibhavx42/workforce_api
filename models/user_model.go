package models


import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Email    string             `json:"location,omitempty" validate:"required"`
	Password string             `json:"title,omitempty" validate:"required"`
}
