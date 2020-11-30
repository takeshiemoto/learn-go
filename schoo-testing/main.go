package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) birthday() {
	p.age++
}

func (p person) birthday2() {
	p.age++
}

func main() {
	terry := &person{
		name: "Terry",
		age:  10,
	}
	terry.birthday()  // ポインタを変更しているので値が書き換わる
	terry.birthday2() // 関数実行時にはコピーが作成されているので値は書き換わらない

	fmt.Printf("%+v\n", terry)
}
