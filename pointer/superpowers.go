package main

import "fmt"

func main() {
	superpowers := &[3]string{"flight", "invisibility", "super strength"}
	fmt.Println(superpowers[0]) // 自動デリファレンス！
	fmt.Println(superpowers[1:2])
}
