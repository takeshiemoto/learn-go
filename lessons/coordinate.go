package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {
	lat := coordinate{
		d: 4,
		m: 35,
		s: 22.2,
		h: 'S',
	}
	long := coordinate{
		d: 127,
		m: 26,
		s: 30.12,
		h: 'E',
	}
	fmt.Println(lat.decimal(), long.decimal())
}
