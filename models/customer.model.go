package models

import (
	"context"

	"github.com/devrodriguez/multitienda-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `bson:"name,omitempty" json:"name,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	PhoneNumber  string             `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	Password     string             `bson:"password,omitempty" json:"password,omitempty"`
	SessionToken string             `bson:"sessionToken,omitempty" json:"sessionToken,omitempty"`
}

func (c Customer) FindOne() (error, Customer) {
	var customer *Customer

	mgClient := *db.GetClient()

	filter := bson.M{"email": c.Email}

	customersRef := mgClient.Database("multitienda").Collection("customers")
	err := customersRef.FindOne(context.TODO(), filter).Decode(&customer)

	if err != nil {
		return err, Customer{}
	}

	return nil, *customer
}

func (c Customer) UpdateToken(token string) (int64, error) {
	mgClient := *db.GetClient()
	filter := bson.M{"email": c.Email}

	customerRef := mgClient.Database("multitienda").Collection("customers")
	updateRes, err := customerRef.UpdateOne(context.TODO(),
		filter,
		bson.D{
			{"$set", bson.D{{"sessionToken", token}}},
		},
	)

	if err != nil {
		return 0, err
	}

	return updateRes.ModifiedCount, nil
}
