package main

import "fmt"

func main() {
	//members := [...]string{
	//	"John",
	//	"Paul",
	//	"George",
	//	"Ringo",
	//}
	//vocals := members[0:2]
	//percussion := members[2:4]
	//fmt.Println(vocals)
	//fmt.Println(percussion)
	//
	//vocals[0] = "Mick"
	//fmt.Println(members)
	//fmt.Println(vocals)
	//
	//teams := []string{"A", "B", "C"}
	//fmt.Println(teams)

	// 0から数えて
	// [0から数えてn番目:1から数えてn番目]
	items := []string{"A", "B", "C", "D", "E"}
	// DとEを取り出して
	myItemsA := items[3:5]
	fmt.Println(myItemsA)
	// BCDを取り出したい
	myItemsB := items[1:4]
	fmt.Println(myItemsB)
}
