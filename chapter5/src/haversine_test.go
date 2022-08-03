package src

import "testing"

func TestHaversine(t *testing.T) {
	tests := []struct {
		p      Point
		q      Point
		distKm float64
	}{
		{
			Point{Lat: 37.983972, Lng: 23.727806}, // Athens
			Point{Lat: 52.366667, Lng: 4.9},       // Amsterdam
			3786.251258825624,
		},
		{
			Point{Lat: 52.366667, Lng: 4.9},       // Amsterdam
			Point{Lat: 52.516667, Lng: 13.388889}, // Berlin
			7557.78,
		},
		{
			Point{Lat: 52.516667, Lng: 13.388889}, // Berlin
			Point{Lat: 37.983972, Lng: 23.727806}, // Athens
			6465,
		},
	}
	for _, input := range tests {
		km := calculateDistance(input.p, input.q)
		if input.distKm != km {
			t.Errorf("fail: wanted %v %v -> %v %v but got %v %v ",
				input.p,
				input.q,
				input.distKm,
				km,
			)
		}
	}
}
