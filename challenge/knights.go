package main

import "fmt"

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v picks up a %v`\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) giv(to *character) {
	if c == nil || to == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v has nothing to give\n", c.name)
		return
	}
	if to.leftHand != nil {
		fmt.Printf("%v`s hands are full`n", to.name)
		return
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v gives %v a %v\n", c.name, to.name, to.leftHand.name)
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v iss carrying nothing", c.name)
	}
	return fmt.Sprintf("%v iss carrying a %v", c.name, c.leftHand.name)
}

func main() {
	artur := &character{
		name: "Artur",
	}
	shrubbery := &item{name: "shrubbery"}
	artur.pickup(shrubbery)

	knigth := &character{
		name: "Knigth",
	}
	artur.giv(knigth)
	fmt.Println(artur)
	fmt.Println(knigth)
}
