package main

import "fmt"

func main() {
	const hoursPerDay = 24
	const days = 28
	const distance = 560000000

	fmt.Println(distance/(days*hoursPerDay), "km/h")
}
