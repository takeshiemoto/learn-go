package main

import (
	"fmt"
)

type Animal interface {
	Run()
}

type Gorilla struct {
	Name string
}

func (g Gorilla) Run() {
	fmt.Println(g.Name, "is Run")
}

func (g Gorilla) Uho() {
	fmt.Println(g.Name, "is Uho")
}

// 空のinterfaceは特定のメソッドを持っていない = 0個のメソッドを持っていれば代入可能！any！
func echo(s interface{}) {
	fmt.Println(s)
}

func main() {
	// GorillaはAnimalIFを満たすので代入が可能
	// Interfaceはメソッドの集まり
	// Gorilla型でも良いがAnimalの方が抽象度が上がる
	var animal Animal = Gorilla{Name: "John"}
	animal.Run()

	// Interfaceから他の型に変換する事をアサーションと呼ぶ
	// 型.(他の型)でアサーションできる
	// 変換先もInterfaceを満たしている必要あり

	// 変換前
	// AnimalだからUhoが呼べない
	// Gorilla型にアサーションする
	g, ok := animal.(Gorilla)
	if ok {
		g.Uho()
	} else {
		fmt.Println("Gorillaへのアサーションに失敗")
	}
}
