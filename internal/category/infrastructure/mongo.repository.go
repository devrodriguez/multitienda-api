package category

import (
	"context"
	"log"
	"time"

	domain "github.com/devrodriguez/multitienda-api/internal/category/domain"
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

// NewMongoRepository return a new implentation
func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (domain.RepositoryContract, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}

	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "category.repository.NewMongoRepository")
	}

	repo.client = client
	return repo, nil
}

func (r *mongoRepository) GetAll() ([]*domain.Category, error) {
	var categories []*domain.Category
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	findOptions := options.Find()
	categoriesRef := r.client.Database(r.database).Collection("categories")
	categoriesCur, err := categoriesRef.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, errors.Wrap(err, "repository.GetAll")
	}

	for categoriesCur.Next(context.TODO()) {
		var category domain.Category

		err := categoriesCur.Decode(&category)
		if err != nil {
			panic(err)
		}

		categories = append(categories, &category)
	}

	return categories, nil
}
