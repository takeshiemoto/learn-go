package main

import "fmt"

func main() {
	members := [...]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	vocals := members[0:2]
	percussion := members[2:4]
	fmt.Println(vocals)
	fmt.Println(percussion)

	vocals[0] = "Mick"
	fmt.Println(members)
	fmt.Println(vocals)

	teams := []string{"A", "B", "C"}
	fmt.Println(teams)

}
