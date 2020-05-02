package models

type GpsPoint struct {
	Lat float64 `bson:"lat,omitempty" json:"lat,omitempty"`
	Lon float64 `bson:"lon,omitempty" json:"lon,omitempty"`
}
