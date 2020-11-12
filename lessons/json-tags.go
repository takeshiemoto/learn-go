package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	point := location{-1.234, 123.1234}
	bytes, _ := json.Marshal(point)

	fmt.Println(string(bytes))
}
