package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"myname" db:"tname" ini:"fname"` //解决其他
	Age  uint8  `json:"myage"`                         //解决(当前端只接收小写字母时，该咋办？)问题
	// class string
}

//序列化
func structMarshal() {
	// p1 := person{Name: "jack", Age: 18, class: "one"}
	p1 := person{Name: "jack", Age: 18}
	value, err := json.Marshal(p1)
	//当结构体中存在非大写开头字段，json是无法接收到的
	//所以value的值中不会出现class字段
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p1)
	fmt.Printf("%v\n", string(value))
}

//反序列化
func structUnmarshal() {
	str := `{"myname":"harry", "myage":26}`
	var p2 person
	err := json.Unmarshal([]byte(str), &p2) //不能传值p2,否则只会更改副本而不会更改原件
	//！！！函数传值等于拷贝！！！,不会对原件起作用
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)
}
func main() {
	// structMarshal()
	structUnmarshal()
}
