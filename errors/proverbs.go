package main

import (
	"fmt"
	"os"
)

// メリット
// エラーを処理するコードに一貫した字下げがある
// エラー処理の繰り返しをすべて丹念に読む代わりにコード全体を読むことが容易
// エラーをインデントするのはGoコミュニティの共通パターン
func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	// return前に必ず実行される
	defer f.Close()

	_, err = fmt.Fprintln(f, "Error are values .")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Don`t just check errors, handle them gracefully .")
	return err
}

func main() {
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
