package main

import "fmt"

func main() {
	teams := []string{"A", "B", "C"}
	// 配列に要素を追加する
	teams = append(teams, "D")
	fmt.Println(teams) // [A B C D]
	// appendは可変長引数を持つ
	teams = append(teams, "E", "F")
	fmt.Println(teams) // [A B C D E F]
}
