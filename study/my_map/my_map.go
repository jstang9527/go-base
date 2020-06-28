package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        //未初始化则不开辟内存空间
	m1 = make(map[string]int, 10) //尽量先估算，避免程序运行时再动态扩容
	m1["a"] = 10
	m1["b"] = 11
	fmt.Println(m1, m1["a"])

	value, ok := m1["b"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}
	//删除item
	delete(m1, "b")
	fmt.Println(m1)
}
