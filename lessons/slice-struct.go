package main

import "fmt"

func main() {
	type location struct {
		name string
		lat  float64
		long float64
	}

	locations := []location{
		{
			name: "A",
			lat:  -4.1234,
			long: 123.12345,
		},
		{
			name: "B",
			lat:  -5.1234,
			long: 133.12345,
		},
	}
	fmt.Println(locations)
}
