package main

import "fmt"

func main() {
	dwarfs := [5]string{"ケレス", "冥王星", "ハウメア", "マケマケ", "エリス"}
	//for i := 0; i < len(dwarfs); i++ {
	//	dwarf := dwarfs[i]
	//	fmt.Println(dwarf)
	//}
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}
}
