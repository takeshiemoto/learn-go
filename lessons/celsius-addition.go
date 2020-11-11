package main

import "fmt"

// 型宣言
type myStr string

// 振る舞いを追加
func (s myStr) sayHello() string {
	return "Hello world"
}

func main() {
	type celsius float64
	var temperature celsius = 20
	fmt.Println(temperature)

	var s myStr
	fmt.Println(s.sayHello())
}
