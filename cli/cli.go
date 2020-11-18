package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
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
		fmt.Printf("%s\t%d\n", item.Name(), item.Size())
	}
}
