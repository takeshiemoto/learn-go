package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	for i := range s {
		s[i] *= 10
	}
	fmt.Println(s)
}
