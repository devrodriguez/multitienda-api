package customer

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	domain "github.com/devrodriguez/multitienda-api/internal/customer/domain"
	"github.com/pkg/errors"
)

type mongoAdapter struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Print(client)
	return client, nil
}

func NewMongoAdapter(mongoURL, database string, timeout int) (domain.CustomerPortIn, error) {
	adapter := &mongoAdapter{
		database: database,
		timeout:  time.Duration(timeout) * time.Second,
	}

	client, err := newMongoClient(mongoURL, timeout)
	if err != nil {
		return nil, errors.Wrap(err, "customer.adapter.NewMongoAdapter")
	}

	adapter.client = client

	return adapter, nil
}

func (a *mongoAdapter) GetAll() ([]*domain.Customer, error) {
	var customers []*domain.Customer
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	findOptions := options.Find()
	colRef := a.client.Database(a.database).Collection("customers")
	colCur, err := colRef.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, errors.Wrap(err, "customer.adapter.GetAll")
	}

	for colCur.Next(context.TODO()) {
		var customer domain.Customer

		err := colCur.Decode(&customer)
		if err != nil {
			panic(err)
		}

		customers = append(customers, &customer)
	}

	return customers, nil
}

func (a *mongoAdapter) Create(customer domain.Customer) error {
	colRef := a.client.Database(a.database).Collection("customers")
	_, err := colRef.InsertOne(context.TODO(), customer)
	if err != nil {
		return err
	}

	return nil
}
