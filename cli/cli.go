package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// flag.BoolはBoolの構造体を返す
	showAll := flag.Bool("a", false, "all files")
	flag.Parse()
	dir := "."
	if len(flag.Args()) > 0 {
		dir = flag.Args()[0]
	}
	items, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("name\tsize")
	for _, item := range items {
		name := item.Name()
		if strings.HasPrefix(name, ".") {
			// 構造体を複合
			if *showAll {
				println(name)
				continue
			}
			continue
		}

		fmt.Printf("%s\t%d\n", item.Name(), item.Size())
	}
}
