package src

import "math"

type Point struct {
	Lat float64
	Lng float64
}

func toRadians(d float64) float64 {
	return d * math.Pi / 180
}

func calculateDistance(cityA, cityB Point) (km float64) {
	latA := toRadians(cityA.Lat)
	latB := toRadians(cityB.Lat)
	lngA := toRadians(cityA.Lng)
	lngB := toRadians(cityB.Lng)
	a := math.Pow(math.Sin((latA-latB)/2), 2) + math.Cos(latA)*math.Cos(latB)*
		math.Pow(math.Sin((lngA-lngB)/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	km = c * 6371
	return km
}
