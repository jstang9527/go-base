package main

import "fmt"

//map和切片的组合
func main() {
	//元素类型为map的切片
	// var s1 = make([]int, 5, 5) int类型切片
	var s1 = make([]map[int]bool, 5, 5) //map类型切片
	s1[0] = make(map[int]bool, 1)
	s1[0][10] = true
	s1[1] = make(map[int]bool, 2)
	s1[1][10] = true
	s1[1][20] = true
	fmt.Println(s1)

	//值为切片的map
	var s2 = make(map[string][]int, 5)
	s2["jk"] = []int{11, 22, 33}
	fmt.Println(s2)
}
