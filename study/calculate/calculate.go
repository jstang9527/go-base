package main

import "fmt"

func main() {
	//位运算
	//5的二进制表示: 101
	//2的二进制表示:  10

	// & 按位与
	fmt.Println(5 & 2) //0
	// | 按位或
	fmt.Println(5 | 2) //111 =>4+2+1=7
	// ^ 按位异或
	fmt.Println(5 ^ 2) //111 =>4+2+1=7
	// << 左移
	fmt.Println(5 << 1) //1010 =>8+0+2+0=10
	// >> 右移
	fmt.Println(5 >> 1) //10 => 2+0=2
}
