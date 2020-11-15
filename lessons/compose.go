package main

func main() {
	type celsius float64
	type temperature struct {
		high, low celsius
	}
	type location struct {
		lat, long float64
	}
	type report struct {
		sol         int
		temperature temperature
		location    location
	}
}
