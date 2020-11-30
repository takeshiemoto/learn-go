package main

import "fmt"

type Square struct {
	width  float64
	height float64
}

func (s Square) Area() float64 {
	return s.width * s.height
}

func (s *Square) Reshape(w float64, h float64) {
	s.width = w
	s.height = h
}

func main() {
	square := Square{3.0, 4.0}
	fmt.Println(square.Area())

}
