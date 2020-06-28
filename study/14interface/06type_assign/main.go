package main

import (
	"fmt"
	"reflect"
)

//类型断言

func singleAssign(a interface{}) {
	if v, ok := a.(string); ok {
		fmt.Printf("是string类型,%#v\n", v)
	} else {
		fmt.Printf("非string类型,%#v\n", v) //v类型零值
	}
}

func assign(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Println("接口是int类型", v)
	case string:
		fmt.Println("接口是string类型", v)
	case bool:
		fmt.Println("接口是布尔类型", v)
	default:
		fmt.Println("无法判定", v)
	}
}

func main() {
	assign(100)
	singleAssign(50)
	var a map[string]interface{}
	a = make(map[string]interface{}, 10)
	a["as"] = [...]int{1, 2, 3, 4}
	fmt.Println(reflect.TypeOf(a))
}
