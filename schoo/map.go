package main

import "fmt"

func main() {
	// nil
	var m map[string]int
	fmt.Printf("%#+v\n", m)

	var m2 = map[string]int{}
	fmt.Println(m2)

	var m3 = make(map[string]int)
	fmt.Println(m3)

	// このパターンおすすめ
	m4 := map[string]int{}
	fmt.Println(m4)

	m5 := map[string]int{"a": 1, "b": 2, "c": 4}
	fmt.Printf("%#+v\n", m5)

	// 取り出す
	v, has := m5["a"]
	if has {
		fmt.Println(v)
	}
}
