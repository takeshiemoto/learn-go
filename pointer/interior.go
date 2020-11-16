package main

import "fmt"

type status struct {
	level             int
	endurance, health int
}

// statusを書き換えるためポインタが必要
func levelUp(s *status) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name   string
	status status
}

func main() {
	player := character{name: "Matthias"}
	levelUp(&player.status)
	fmt.Printf("%+v\n", player.status)
}
