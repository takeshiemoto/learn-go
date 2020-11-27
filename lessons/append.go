package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: 長さ %v, 容量 %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	//s1 := []string{"A", "B", "C"}
	//dump("s1", s1)
	//s2 := append(s1, "D")
	//dump("s2", s2)
	//s3 := append(s2, "E")
	//dump("s3", s3)
	//
	//base := []string{"a", "b", "c", "d", "e"}
	//
	//b1 := base[0:4:4]
	//dump("b1", b1)

	var x []string
	dump("x", x)
	y := []string{}
	dump("y", y)

}
