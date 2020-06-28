package main

import (
	"fmt"
	"strings"
)

//闭包是什么？
//闭包是一个函数，这个函数包含了他外部作用域变量
func myFunc1(f func()) {
	fmt.Println("myFunc1")
	f()
}
func myFunc2(x, y int) {
	fmt.Println("myFunc2")
	fmt.Println(x + y)
}

//为让f2可以被f1调用，对f2进行包装
func myFunc3(f func(int, int), m, n int) func() {
	fmt.Println("myFunc3")
	return func() {
		f(m, n) //执行f2
	}
}
func main() {
	ret := myFunc3(myFunc2, 30, 30)
	myFunc1(ret)

	//后缀例子
	fmt.Println("=====后缀例子=====")
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("meitu"))
	fmt.Println(jpgFunc("go.jpg"))
	fmt.Println(txtFunc("dushu"))
	fmt.Println(txtFunc("yufu.txt"))
}

//例子，后缀名
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
