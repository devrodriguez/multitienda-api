package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Email       string             `bson:"email,omitempty" json:"email,omitempty"`
	PhoneNumber string             `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	Password    string             `bson:"password,omitempty" json:"password,omitempty"`
}
