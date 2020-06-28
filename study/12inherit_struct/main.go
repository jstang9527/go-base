package main

import "fmt"

//继承

//动物类
type animal struct {
	name string
}

//动物通用方法
func (a animal) move() {
	fmt.Printf("%s会移动\n", a.name)
}

//狗类
type dog struct {
	id uint8
	animal
}

//狗通用方法
func (d dog) run() {
	fmt.Printf("%s xiuxiu~ ~\n", d.name)
}

func main() {
	d1 := dog{
		id:     86,
		animal: animal{name: "jack"},
	}
	fmt.Println(d1)
	d1.move()
	d1.run()
}
