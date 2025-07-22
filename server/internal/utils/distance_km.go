package utils

import (
	"log"
	"math"
)

// Earth radius in kilometers
const EarthRadius = 6371

// Degrees to radians
func degToRad(deg float64) float64 {
	return deg * math.Pi / 180
}

// Haversine distance in kilometers
func DistanceInKM(lat1, lon1, lat2, lon2 float64) float64 {
	log.Printf("Calculating distance between (%f, %f) and (%f, %f)", lat1, lon1, lat2, lon2)
	dLat := degToRad(lat2 - lat1)
	dLon := degToRad(lon2 - lon1)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(degToRad(lat1))*math.Cos(degToRad(lat2))*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return EarthRadius * c
}
