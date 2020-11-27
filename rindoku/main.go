package main

import "fmt"

func Sum(a []float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}

func main() {
	s := []float64{3.5, 5.5, 2.5}
	sum := Sum(s)
	fmt.Println(sum) // 11.5
}
