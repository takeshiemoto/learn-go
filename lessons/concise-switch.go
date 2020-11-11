package main

import "fmt"

func main() {
	fmt.Println("洞窟の入口だ。東に進む道もある ")

	const command = "go inside"

	switch command {
	case "go east":
		fmt.Println("きみは、さらに山を登る。 ")
	case "enter cave", "go inside": // カンマで区切った複数の候補
		fmt.Println("きみは、薄暗い洞窟の中にいる。 ")
	case "read sign":
		fmt.Println("「未成年立入禁止」と書いてある。 ")
	default:
		fmt.Println("なんだかよくわからない。 ")
	}
}
