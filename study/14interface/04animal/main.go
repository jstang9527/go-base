package main

import (
	"fmt"
)

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet uint8
}
type bird struct {
	name string
	feet uint8
}

func (c cat) move() {
	fmt.Println("跑", c.feet)
}
func (b bird) move() {
	fmt.Println("飞", b.feet)
}
func (c cat) eat(s string) {
	fmt.Println("吃鱼", c.name)
}
func (b bird) eat(s string) {
	fmt.Println("吃虫", b.name)
}
func action(a animal, s string) {
	a.move()
	a.eat(s)
}
func main() {
	var a1 animal
	c1 := cat{
		name: "tom",
		feet: 4,
	}
	a1 = c1
	action(a1, "asd")
}
