package main

import "fmt"

func main() {
	type person struct {
		name, superpwer string
		age             int
	}

	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	fmt.Println(timmy)

	// 構造体はデリファレンス不要で値の書き換えが可能
	timmy.superpwer = "flying"
	fmt.Println("%+v\n", timmy)
}
