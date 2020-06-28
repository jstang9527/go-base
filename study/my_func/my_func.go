package main

import "fmt"

func test() (x int) {
	return 5
}

//函数作为参数和返回值
func myFunc(x int, f func() int) func(int, int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", f)
	return func(a int, b int) int {
		return a + b
	}
}

//匿名函数
func father(a int, b int) int {
	return func() int {
		return a + b
	}()
}

func main() {
	//高阶函数
	// fmt.Printf("%T\n", test) //func() int
	// ff := myFunc(99, test)
	// ret := ff(3, 4)
	// fmt.Printf("%d\n", ret)

	//匿名函数
	var f1 = func(x, y int) int { return x + y }
	fmt.Println(f1(33, 66))
	//匿名函数在函数的使用(真正解决在函数中无法声明命名函数)
	fmt.Println(father(10, 20))
	//立即执行函数(如函数只使用一次推荐该方式)
	func(x, y int) { fmt.Printf("立即执行%d+%d=%d\n", x, y, x+y) }(100, 200)
}
