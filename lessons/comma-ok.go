package main

import "fmt"

func main() {
	members := map[string]int{
		"Syuto":     23,
		"Makamura":  7,
		"Yanagita":  9,
		"Despaigne": 54,
		"Kurihara":  32,
	}

	if member, ok := members["Syuto"]; ok {
		fmt.Println("OK")
		fmt.Println(ok)
		fmt.Println(member)
	} else {
		fmt.Println("No")
		fmt.Println(member)
		fmt.Println(ok)
	}
}
