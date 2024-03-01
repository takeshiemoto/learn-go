package main

import (
	"fmt"
	"strconv"
)

func main() {
	//run0()
	run1()
}

func run0() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	var product = a * b
	if product%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func run1() {
	var input string
	fmt.Scanf("%s", &input)

	a, _ := strconv.Atoi(string(input[0]))
	b, _ := strconv.Atoi(string(input[1]))
	c, _ := strconv.Atoi(string(input[2]))

	fmt.Println(a + b + c)
}
