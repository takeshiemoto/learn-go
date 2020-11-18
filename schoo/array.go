package main

import "fmt"

func main() {
	n := [...]int{1, 2, 3, 4, 5}
	n2 := n[1:3]
	fmt.Println(len(n2), cap(n2))

	// スライスではappendで容量を増やすことができる
	n2 = append(n2, 6, 7)
	fmt.Println(len(n2), cap(n2))

	n2 = append(n2, 6, 7)
	fmt.Println(len(n2), cap(n2))

	n2 = append(n2, 8, 9)
	fmt.Println(len(n2), cap(n2))

	/* 初期化方法 */
	var i1 []int
	fmt.Println(i1)

	i2 := []int{}
	fmt.Printf("%#v\n", i2)
	fmt.Println(i2)

	i3 := make([]int, 2, 3)
	fmt.Println(i3)

	i4 := []int{1, 2, 3, 4}
	fmt.Println(i4)

}
