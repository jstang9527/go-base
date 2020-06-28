package main

import "fmt"

// func main() {
// 	//99乘法表
// 	for i := 1; i < 10; i++ {
// 		for j := 1; j <= i; j++ {
// 			fmt.Printf("%v x %v=%v  ", i, j, i*j)
// 		}
// 		println()
// 	}
// }
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	a, b := 1, 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
