package main

import (
	"fmt"
)

type speaker interface {
	speak() //只要实现了这个方法的变量都是speaker的类型
}

type person struct{}
type dog struct{}
type cat struct{}

func (p *person) speak() {
	fmt.Println("啊啊啊")
}
func (c *cat) speak() {
	fmt.Println("喵喵喵")
}
func (d *dog) speak() {
	fmt.Println("汪汪汪")
}

func hit(s speaker) {
	s.speak()
}
func main() {
	var p person
	var c cat
	var d dog
	hit(&p)
	hit(&c)
	hit(&d)
}
