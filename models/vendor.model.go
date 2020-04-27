package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vendor struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Address     string             `bson:"address,omitempty" json:"address,omitempty"`
	PhoneNumber string             `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
}
