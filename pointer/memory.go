package main

import "fmt"

func main() {
	answer := 42
	// アドレス演算子(%)はメモリのアドレスを提供する
	fmt.Println(&answer) // 0xc0000180e8

	// (*)はデリファレンス(逆参照)
	address := &answer
	fmt.Println(*address)
}
