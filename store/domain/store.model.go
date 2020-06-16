package store

import (
	category "github.com/devrodriguez/multitienda-api/category/domain"
	"github.com/devrodriguez/multitienda-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name           string             `bson:"name,omitempty" json:"name,omitempty"`
	Category       category.Category  `bson:"category,omitempty" json:"category,omitempty"`
	Description    string             `bson:"description,omitempty" json:"description,omitempty"`
	UrlImages      []string           `bson:"urlImages,omitempty" json:"urlImages,omitempty"`
	String64Images []string           `bson:"string64Images,omitempty" json:"string64Images,omitempty"`
	Address        string             `bson:"address,omitempty" json:"address,omitempty"`
	PhoneNumber    string             `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	Status         string             `bson:"status,omitempty" json:"status,omitempty"`
	GeoLocation    models.GpsPoint    `bson:"geolocation,omitempty" json:"geolocation,omitempty"`
}
