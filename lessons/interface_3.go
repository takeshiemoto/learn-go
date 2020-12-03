package main

func main() {
	// panic
	//i := interface{}("Hello")
	//s := i.(int)
	//fmt.Println(s)

	// 第２引数を受け取ればpanicにならない
	// この時、値はアサーションした型の0値となる
	//i := interface{}("Hello")
	//n, ok := i.(int)
	//fmt.Println(n, ok) // 0 false

	// switchを使って複数の方をチェックする
	//i := interface{}("Hello")
	//switch i.(type) {
	//case string:
	//	s := i.(string)
	//	fmt.Println("i is string", s)
	//case bool:
	//	b := i.(bool)
	//	fmt.Println("i is bool", b)
	//case int:
	//	n := i.(int)
	//	fmt.Println("i is int", n)
	//}

	// 構造体の変換も可能
	//type A struct{}
	//i := interface{}(A{})
	//_, ok := i.(A)
	//fmt.Println(ok) // true
}
