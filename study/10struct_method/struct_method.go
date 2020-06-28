package main

import "fmt"

//Dog 狗
type Dog struct {
	name string
}

//Person 人
type Person struct {
	name string
	age  uint8
}

//方法：作用于特定类型的函数
//func (接收者 类型) 函数名(参数) (结果参数){}
func (d Dog) run() {
	fmt.Println(d.name, "xiuxiu...")
}
func (p Person) run() {
	fmt.Println(p.name, "驾~驾~~~")
}
func (p Person) growthValue() { //值类型则拷贝++,原件无改变
	p.age++
}
func (p *Person) growthPoint() { //直接操作原件，原件改变, 直接修改接收者的值
	p.age++
}

//构造函数
func newDog(name string) *Dog {
	return &Dog{name: name}
}
func newPerson(name string, age uint8) *Person {
	return &Person{name: name, age: age}
}

func main() {
	// myDog := Dog{name: "天庭哮天犬"}
	fmt.Println("来将可留姓名！")
	// myDog.run()
	// hero := Person{name: "常山赵子龙"}
	// hero.run()
	m1 := newPerson("吾乃零陵上将军邢道荣!", 26)
	m1.run()
	fmt.Println(m1.age)
	m1.growthValue()
	fmt.Println(m1.age)
	m1.growthPoint()
	fmt.Println(m1.age)
}
