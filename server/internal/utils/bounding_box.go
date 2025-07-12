package utils

import "math"

type BoundingBox struct {
	MinLat float64
	MaxLat float64
	MinLong float64
	MaxLong float64
}

// GetLatLonBoundsMeters returns latitude/longitude bounds given a center and radius in meters
func GetLatLongBoundsMeters(lat, long, radiusMeters float64) BoundingBox {
	// Convert radius to kilometers
	radiusKm := radiusMeters / 1000.0

	// 1 degree latitude ≈ 111 km
	latDelta := radiusKm / 111.0

	// Adjust longitude delta by latitude
	longDelta := radiusKm / (111.320 * math.Cos(lat*math.Pi/180))

	return BoundingBox{
		MinLat: lat - latDelta,
		MaxLat: lat + latDelta,
		MinLong: long - longDelta,
		MaxLong: long + longDelta,
	}
}
