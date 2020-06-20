package category

import (
	"context"
	"time"

	"github.com/devrodriguez/multitienda-api/internal/shared"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type adapterOut struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

// NewMongoRepository return a new implentation
func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (PortOut, error) {
	repo := &adapterOut{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}

	client, err := shared.NewMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "category.repository.NewMongoRepository")
	}

	repo.client = client
	return repo, nil
}

func (r *adapterOut) GetAllDB() ([]*Category, error) {
	var categories []*Category
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	findOptions := options.Find()
	categoriesRef := r.client.Database(r.database).Collection("categories")
	categoriesCur, err := categoriesRef.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, errors.Wrap(err, "repository.GetAllDB")
	}

	for categoriesCur.Next(context.TODO()) {
		var category Category

		err := categoriesCur.Decode(&category)
		if err != nil {
			panic(err)
		}

		categories = append(categories, &category)
	}

	return categories, nil
}
