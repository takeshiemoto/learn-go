package main

import "fmt"

const PiTyped float64 = 3.14
const PiUntyped = 3.14

const AddedValue = 100

func main() {
	fmt.Println(Add(PiUntyped))
}

func Add(value int) int {
	return value + AddedValue
}

func SayHello(s string) string {
	return "Hello " + s
}
