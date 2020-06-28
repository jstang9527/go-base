package main

import "fmt"

type address struct {
	province string
	city     string
}
type person struct {
	name string
	age  uint8
	addr address //嵌套结构体
}
type company struct {
	name    string
	address //匿名嵌套结构体
}

func main() {
	p1 := person{
		name: "jack",
		age:  18,
		addr: address{
			province: "广东",
			city:     "广州",
		},
	}
	c1 := company{
		name: "huawei",
		address: address{
			province: "上海",
			city:     "浦东",
		},
	}
	fmt.Println(p1, c1)
	fmt.Println(p1.addr.city)
	fmt.Println(c1.city) //变量先从自身结构体找，然后再去匿名结构体找
	//注意：如果匿名嵌套出现相同的字段，则只能这样写:ci.address.city
}
