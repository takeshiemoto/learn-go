package main

import "fmt"

// 関数は値渡しの為namesの変更はできていない!
func terraform(names [3]string) {
	for i := range names {
		names[i] = "New " + names[i]
	}
}

func main() {
	names := [...]string{"A", "B", "C"}
	terraform(names)
	fmt.Println(names)
}
