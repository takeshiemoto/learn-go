package main

import (
	"fmt"
)

type Member struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	/* 構造体を宣言 */
	var g Member
	g.FirstName = "David"
	g.LastName = "Gilmour"
	fmt.Printf("%+v\n", g)

	/* 構造体を複合リテラルで初期化する */
	var n = Member{
		FirstName: "Nick",
		LastName:  "Mason",
	}
	fmt.Printf("%+v\n", n)

	/* フィールド名は省略可能 */
	var r = Member{"Richard", "Wright"}
	fmt.Printf("%+v\n", r)

	/* フィールドに値が無い場合はゼロ値が入る */
	var v = Member{
		FirstName: "Roger",
	}
	fmt.Printf("%+v\n", v)
	v.LastName = "Waters"
	fmt.Printf("%+v\n", v)

	// 構造体はコピー
	s := g
	s.FirstName = "Syd"
	s.LastName = "Barrett"
	fmt.Printf("%+v\n", g)
	fmt.Printf("%+v\n", s)

	floyd := []Member{
		{
			FirstName: "David",
			LastName:  "Gilmour",
		},

		{
			FirstName: "Nick",
			LastName:  "Mason",
		},
		{
			FirstName: "Richard",
			LastName:  "Wright",
		},
		{
			FirstName: "Roger",
			LastName:  "Waters",
		},
	}
	fmt.Printf("%+v\n", floyd)

}
