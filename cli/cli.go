package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	items, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("name\tsize")
	for _, item := range items {
		fmt.Printf("%s\t%d\n", item.Name(), item.Size())
	}
}
