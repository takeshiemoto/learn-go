package main

import "fmt"

func myFunc() string {
	return "Hello world"
}

func main() {
	if a := myFunc(); a != "Hello my world" {
		fmt.Println("NG")
		return
	}
	fmt.Println("OK")
}
