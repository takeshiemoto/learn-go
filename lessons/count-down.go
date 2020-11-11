package main

import "fmt"

func main() {
	//var count = 10
	//for count > 0 {
	//	fmt.Println(count)
	//	time.Sleep(time.Second)
	//	count--
	//}
	//fmt.Println("Liftoff!!")

	for count := 10; count < 0; count-- {
		fmt.Println(count)
	}
}
