package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

// 関数は値渡しだがポインタを使うことで参照を直接書き換えられる！
func birthday(p *person) {
	p.age++
}

// ポインタのレシーバでも同様
func (p *person) birthday() {
	p.age++
}

func main() {
	paul := person{
		name:       "Paul",
		superpower: "Flying",
		age:        29,
	}

	george := person{
		name:       "George",
		superpower: "Flying",
		age:        22,
	}

	george.birthday()

	fmt.Println(george) // {George Flying 23}

	birthday(&paul)

	fmt.Println(paul) // {Paul Flying 30}
}
