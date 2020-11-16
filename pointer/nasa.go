package main

import "fmt"

func main() {
	var administrator *string // administratorポインタ
	john := "John"
	administrator = &john // administrator = "Jhon"
	fmt.Println(*administrator)

	paul := "Paul"
	administrator = &paul
	fmt.Println(*administrator) // Paul

	// デリファレンスして書き換える
	*administrator = "George" // *administerに格納された値を書き換える
	fmt.Println(paul)         // PaulがGeorgeになっている！
}
