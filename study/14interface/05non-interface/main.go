package main

import "fmt"

func show(value interface{}) { //可以传任意类型的变量
	fmt.Printf("%#v\n", value)
}

func main() {
	var person map[string]interface{}
	person = make(map[string]interface{}, 10)
	person["name"] = "golang"
	person["age"] = 18
	person["isVIP"] = false
	person["hobby"] = [...]string{"蹦迪", "泡吧", "夜总会"}
	show(18)
	show("F117")
	show(false)
	show(person)
}
