package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temp := temperature["Earth"]
	fmt.Println(temp)

	temperature["Earth"] = 16
	fmt.Println(temperature)

	if moon, ok := temperature["Moon"]; ok {
		fmt.Println(moon)
	} else {
		fmt.Println("?")
	}
}
