package main

import "fmt"

func main() {
	// 型のある定数
	const n int = 100
	// 型の無い定数
	const m = 100
	// 定数式の使用
	const s = "Hello world" + "世界"
	const n2 = 100000000000000000 / 100000000000000000
	var m2 = n2 // 1
	fmt.Println(m2)
}
