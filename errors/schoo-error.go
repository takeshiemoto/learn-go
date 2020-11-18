package main

import (
	"log"
	"os"
)

func main() {
	// GoはTyr-Catchが無い
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	f.Close()
}
