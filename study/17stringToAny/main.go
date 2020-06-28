package main

import (
	"fmt"
	"strconv"
)

func main() {

	//1.字符串转int
	str := "97"
	sInt, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("%T, %#v\n", sInt, sInt)
	//2.字符串转int
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%T, %#v\n", retInt, retInt)

	//3.int转字符串
	i := uint8(97)
	iStr := fmt.Sprintf("%d", i)
	fmt.Printf("%T, %#v\n", iStr, iStr)
	//4.int转字符串
	i2 := 97
	ret := strconv.Itoa(i2)
	fmt.Printf("%T, %#v\n", ret, ret)

	//5.string to bool
	boolStr := "true"
	b, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%T, %#v\n", b, b)

	//6.string to float
	floatStr := "3.14"
	f, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%T, %#v\n", f, f)
}
