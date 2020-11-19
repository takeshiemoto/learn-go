package main

import "fmt"

func main() {
	// IF内で宣言したnameはifの中だけでしか利用できない
	if name := "John"; name == "John" {
		fmt.Println(name)
	}

	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		i++
	}
}
