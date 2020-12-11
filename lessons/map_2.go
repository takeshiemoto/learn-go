package main

import "fmt"

func main() {
	var m1 = map[string]int{
		"John":   0,
		"Paul":   10,
		"George": 20,
		"Ringo":  30,
	}

	// これでは存在しないからゼロ値を返しているのか
	// Johnの値が0なのか判断できない
	fmt.Println(m1["John"]) // 0

	// comma, ok idiom.
	v, ok := m1["John"]
	fmt.Println(ok, v)

	v, ok = m1["Nick"]
	fmt.Println(ok, v)
}
