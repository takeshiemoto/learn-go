package main

import "fmt"

func main() {
	temperature := []float64{
		-28.0,
		32.0,
		-31.0,
		-29.0,
		-23.0,
		-29.0,
		-28.0,
		-33.0,
	}

	// [float64: ture]のMap
	set := make(map[float64]bool)

	for _, t := range temperature {
		// 重複が消される
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("セットのメンバ")
	}

	fmt.Println(set)
}
