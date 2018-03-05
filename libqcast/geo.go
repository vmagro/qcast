package libqcast

import (
	"fmt"
	"math"
)

// Location represents a latitude / longitude pair
type Location struct {
	Lat float64
	Lng float64
}

// NewLocation creates a new Location struct with the given latitude and longitude
func NewLocation(lat, lng float64) *Location {
	return &Location{
		Lat: lat,
		Lng: lng,
	}
}

// DistanceTo computes how far away in meters another point is using the Haversine formula.
func (l *Location) DistanceTo(other *Location) float64 {
	// a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
	// c = 2 ⋅ atan2( √a, √(1−a) )
	// d = R ⋅ c
	phi1 := l.Lat * math.Pi / 180.0
	phi2 := other.Lat * math.Pi / 180.0
	dPhi := phi2 - phi1
	dLambda := (other.Lng - l.Lng) * math.Pi / 180.0
	a := math.Pow(math.Sin(dPhi/2), 2) + math.Cos(phi1)*math.Cos(phi2)*math.Pow(math.Sin(dLambda/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Earth's radius
	R := 637100.0
	dist := R * c
	return dist
}

// String returns a human-readable string for this location
func (l *Location) String() string {
	return fmt.Sprintf("Location:%f,%f", l.Lat, l.Lng)
}
