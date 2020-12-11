package main

import "fmt"

func main() {

	// 重複のあるint型のスライス
	numbers := []int{1, 3, 3, 5, 3, 2, 3, 2}

	// ブール値を値に持つMapを作成
	set := make(map[int]bool)

	for _, n := range numbers {
		// 同一のキーは存在できない為
		set[n] = true
	}

	fmt.Println(set)

	// Mapは順序が任意なのでスライスに戻す
	unique := make([]int, 0, len(set))
	for n2 := range set {
		unique = append(unique, n2)
	}

	fmt.Println(unique)
}
