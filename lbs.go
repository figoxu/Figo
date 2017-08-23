package Figo

import "github.com/kellydunn/golang-geo"

type GeoPoint struct {
	Lat float64
	Lng float64
}

func NewGeoPoint(lat, lng float64) GeoPoint {
	return GeoPoint{
		Lat: lat,
		Lng: lng,
	}
}

func GeoDistance(from, to GeoPoint) float64 {
	return geo.NewPoint(from.Lat, from.Lng).GreatCircleDistance(geo.NewPoint(to.Lat, to.Lng)) * 1000
}
