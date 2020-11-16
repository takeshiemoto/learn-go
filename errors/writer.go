package main

import (
	"fmt"
	"io"
	"os"
)

type saveWriter struct {
	w   io.Writer
	err error // ここに最初のエラーをストアする
}

func (sw *saveWriter) writeln(s string) {
	if sw.err != nil {
		return // 既にエラーが発生していたら書き込みをスキップ
	}
	// 1行を書き、もしエラーがあればストアする
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbs2(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := saveWriter{w: f}
	sw.writeln("Errors aer values .")
	sw.writeln("Dont`t just check errors, handle them gracefully .")

	// エラーが発生したらそのエラーを返す
	return sw.err
}
