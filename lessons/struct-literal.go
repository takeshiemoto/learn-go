package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}
	opportunity := location{lat: 1.12345, long: 354.1234}
	fmt.Println(opportunity)
}
