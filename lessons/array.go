package main

import "fmt"

func main() {
	var planets [8]string
	planets[0] = "彗星"
	planets[1] = "金星"
	planets[2] = "地球"
	earth := planets[2]
	fmt.Println(len(planets))
	fmt.Println(earth)
}
