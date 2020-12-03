package main

import "fmt"

func check(i interface{}) {
	// 引数で受け取った型がstring型かどうかアサーションしている
	// アサーション結果
	// a 変換された値
	// b 変換が成功したか否かのbool
	a, b := i.(string)
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	check("Hello")
}
