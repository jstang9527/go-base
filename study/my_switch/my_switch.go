package main

import (
	"fmt"
)

func main() {
	switch n := 3; n {
	case 1:
		fmt.Println("冠军")
	case 2:
		fmt.Println("亚军")
	case 3:
		fmt.Println("季军")
	default:
		fmt.Println("感谢参与")
	}

	//使用表达式
	score := 70
	switch {
	case score >= 80:
		fmt.Print("优秀")
		fallthrough //当满足该条件后可执行下一个case, 将被舍弃, 尽量不用
	case score >= 60 && score < 80:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

}
