package main

import "fmt"

func fp(xp *int) {
	*xp = 10
}

func f(x int) {
	x = 10
}

func main() {
	n := 100
	fp(&n)
	fmt.Println("n", n)

	x := 100
	f(x)
	fmt.Println("x", x)
}
