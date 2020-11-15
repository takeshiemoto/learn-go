package main

import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pwe", int(1))
}

func main() {
	t = martian{}
	fmt.Println(t.talk())
	t = laser(3)
	fmt.Println(t.talk())
}
