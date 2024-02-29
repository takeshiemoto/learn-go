package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	var product = a * b
	if product%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
