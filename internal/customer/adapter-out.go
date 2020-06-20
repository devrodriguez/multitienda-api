package customer

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/devrodriguez/multitienda-api/internal/shared"
	"github.com/pkg/errors"
)

type AdapterOut struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func NewAdapterOut(mongoURL, database string, timeout int) (PortOut, error) {
	adapter := &AdapterOut{
		database: database,
		timeout:  time.Duration(timeout) * time.Second,
	}

	client, err := shared.NewMongoClient(mongoURL, timeout)
	if err != nil {
		return nil, errors.Wrap(err, "customer.adapter.NewMongoAdapter")
	}

	adapter.client = client

	return adapter, nil
}

func (adapter *AdapterOut) GetAllDB() ([]*Customer, error) {
	var customers []*Customer
	ctx, cancel := context.WithTimeout(context.Background(), adapter.timeout)
	defer cancel()

	findOptions := options.Find()
	colRef := adapter.client.Database(adapter.database).Collection("customers")
	colCur, err := colRef.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, errors.Wrap(err, "customer.adapter.GetAll")
	}

	for colCur.Next(context.TODO()) {
		var customer Customer

		err := colCur.Decode(&customer)
		if err != nil {
			panic(err)
		}

		customers = append(customers, &customer)
	}

	return customers, nil
}

func (adapter *AdapterOut) CreateDB(customer Customer) error {
	colRef := adapter.client.Database(adapter.database).Collection("customers")
	_, err := colRef.InsertOne(context.TODO(), customer)
	if err != nil {
		return err
	}

	return nil
}
