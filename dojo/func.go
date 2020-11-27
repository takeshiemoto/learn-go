package main

import "fmt"

// 何が戻り地になるのかわかりにくい
// 極力使わない
func swap(x, y int) (x2, y2 int) {
	y2, x2 = x, y
	return
}

func main() {
	// リテラル
	// 書き下す
	// 無名関数 = クロージャ
	// int stringと同じ用に扱える（ファーストクラス）
	msg := "Hello world"
	func() {
		fmt.Println(msg)
	}()

	fs := make([]func() string, 2)
	fs[0] = func() string {
		return "Hoge"
	}
	fs[1] = func() string {
		return "Fuga"
	}
	for _, f := range fs {
		fmt.Println(f())
	}
}
