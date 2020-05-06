package models

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name           string             `bson:"name,omitempty" json:"name,omitempty"`
	Category       Category           `bson:"category,omitempty" json:"category,omitempty"`
	Description    string             `bson:"description,omitempty" json:"description,omitempty"`
	UrlImages      []string           `bson:"urlImages,omitempty" json:"urlImages,omitempty"`
	String64Images []string           `bson:"string64Images,omitempty" json:"string64Images,omitempty"`
	Address        string             `bson:"address,omitempty" json:"address,omitempty"`
	PhoneNumber    string             `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	Status         string             `bson:"status,omitempty" json:"status,omitempty"`
	GeoLocation    GpsPoint           `bson:"geolocation,omitempty" json:"geolocation,omitempty"`
}

func (s *Store) GetCoordinates() error {
	var key = "AIzaSyBnh1nBPOMEml3EbvkzGn0c-LYEKMfKhfE"
	var gpsPoint GpsPoint
	var apiRes GoogleApiCoord

	res, err := http.Get("https://maps.googleapis.com/maps/api/geocode/json?address=" + url.QueryEscape(s.Address) + "&key=" + key)
	if err != nil {
		return err
	}

	log.Println(s.Address)

	json.NewDecoder(res.Body).Decode(&apiRes)

	if err != nil {
		return err
	}
	gpsPoint.Lat = apiRes.Results[0].Geometry.Location.Lat
	gpsPoint.Lon = apiRes.Results[0].Geometry.Location.Lng

	s.GeoLocation = gpsPoint

	log.Println(s)

	return nil
}
