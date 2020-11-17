package main

import "fmt"

type Gorilla struct {
	Name string
}

type Animal string

func (a Animal) Uho() {
	fmt.Println(a)
}

// 変数を宣言しなくても良いがあまり使わない
func (Gorilla) Like() string {
	return "Banana"
}

// gは値渡し！渡される値がコピーされている
//func (g Gorilla) SetName(name string) {
//	g.Name = name
//}

// 一般的は構造体で定義する
func (g *Gorilla) SetName(name string) {
	g.Name = name
}

func (g Gorilla) Uho() {
	fmt.Println(g.Name, g.Banana())
}

func (g Gorilla) Banana() string {
	return "Banana"
}

func main() {
	g := Gorilla{Name: "John"}
	g.SetName("Paul")
	fmt.Println(g.Name)

	// キャストしないとUhoは呼べない
	a := Animal("Gorilla")
	a.Uho()
}
