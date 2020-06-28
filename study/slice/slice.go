package main

import (
	"fmt"
	"sort"
)

func main() {
	//append
	s1 := make([]string, 1)
	s1[0] = "oo"
	s2 := []string{"aa", "vv", "bb", "zz"}
	s1 = append(s1, s2...) //...表示将aa、vv、bb拆开
	fmt.Println(s1)

	//将s2的bb删除
	s2 = append(s2[:2], s2[3:]...)
	fmt.Println(s2)
	fmt.Println(cap(s2))

	//切片排序
	var a1 = [...]int{7, 1, 9, 3, 5}
	sort.Ints(a1[:]) //传入切片
	fmt.Println(a1)

	//copy
	var aa = a1[:]
	var a2 = make([]int, 5, 5)
	copy(a2, aa)
	a2[0] = 9
	fmt.Println(&aa[2])
	fmt.Println(&a1[2])
}
