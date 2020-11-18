package main

import "fmt"

var s1 = struct {
	name string
	age  int
}{name: "Person", age: 22}

// 再利用可能になる
type Person struct {
	name string
	age  int
}

func main() {
	s2 := Person{
		name: "John",
		age:  31,
	}
	fmt.Println(s1)
	fmt.Println(s2)
}
