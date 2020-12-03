package main

import "fmt"

type I1 struct {
}

type I2 struct {
}

func (b I2) Hello(name string) {
	fmt.Println("Hello", name)
}

type I interface {
	Hello(name string)
}

func main() {
	// I1構造体はHelloメソッドを持っていない = Iを満たしていない
	// interfaceはanyみたいなもの
	// とりあえずI1型の構造体を定義している
	a := interface{}(I1{})
	_, ok := a.(I)
	fmt.Println(ok)

	// I2構造体はHelloをメソッドを持っているのでIを満たす
	b := interface{}(I2{})
	_, ok = b.(I)
	fmt.Println(ok)
}
