package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, columns = 9, 9

// 多くのGoパッケージはエラーを変数で宣言してエクスポートしている
// 規約により、エラーメッセージを代入する変数はErrで始める
var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

type Grid [rows][columns]int8 // 配列

func inBounds(row, column int) bool {
	if row > 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func (g *Grid) Set(row, column int, digit int8) error {
	if inBounds(row, column) {
		return errors.New("out of bounds")
	}
	g[row][column] = digit
	return nil
}

func main() {
	var g Grid
	err := g.Set(10, 0, 5)
	// エラーメッセージを変数定義することで
	// switch文で分岐できる
	if err != nil {
		// errors.newはポインタ
		// テキスト比較ではなくメモリアドレスを比較している！
		switch err {
		case ErrBounds, ErrDigit:
			fmt.Println("Error!!")
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}
}
