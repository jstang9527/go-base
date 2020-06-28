package main

import "fmt"

func main() {
	//goto(跳转到指定标签),尽量不用
	for i := 0; i < 5; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto breakTag
			}
			fmt.Printf("%v-%c\t", i, j)
		}
		fmt.Println()
	}
breakTag:
	fmt.Println("\n结束全部for循环")

	//推荐此种方式
	var flag = false
	for i := 0; i < 5; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break
			}
			fmt.Printf("%v-%c\t", i, j)
		}
		fmt.Println()
		if flag {
			break
		}
	}
}
