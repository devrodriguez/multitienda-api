package store

import (
	"context"
	"log"
	"time"

	domain "github.com/devrodriguez/multitienda-api/store/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
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

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (RepositoryContract, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}

	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "store.repository.NewMongoRepository")
	}

	repo.client = client
	return repo, nil
}

func (r *mongoRepository) GetAllStores() ([]*domain.Store, error) {
	var categories []*domain.Store
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	findOptions := options.Find()
	categoriesRef := r.client.Database(r.database).Collection("stores")
	categoriesCur, err := categoriesRef.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, errors.Wrap(err, "store.repository.GetAll")
	}

	for categoriesCur.Next(context.TODO()) {
		var category domain.Store

		err := categoriesCur.Decode(&category)
		if err != nil {
			panic(err)
		}

		categories = append(categories, &category)
	}

	return categories, nil
}
