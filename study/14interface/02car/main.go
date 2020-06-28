package main

import "fmt"

type car interface {
	run()
}

type benz struct {
	brand string
}
type bmw struct {
	brand string
}

func (b benz) run() {
	fmt.Println("速度120km/h", b.brand)
}
func (b bmw) run() {
	fmt.Println("速度110km/h", b.brand)
}

//Run X
func drive(c car) {
	c.run()
}

func main() {
	b1 := benz{brand: "奔驰"}
	b2 := bmw{brand: "宝马"}
	drive(b1)
	drive(b2)
}
