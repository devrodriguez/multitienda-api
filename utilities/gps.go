package utilities

import (
	"github.com/devrodriguez/multitienda-api/models"
	"github.com/umahmood/haversine"
)

func HaversineFormule(org, des models.GpsPoint) (float64, float64) {
	var origen = haversine.Coord{org.Lat, org.Lon}
	var destino = haversine.Coord{des.Lat, des.Lon}
	return haversine.Distance(origen, destino)
}
