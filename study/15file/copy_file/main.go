package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile() {
	src, err := os.Open("../log.txt")
	defer src.Close()
	if err != nil {
		fmt.Println("打开源文件发生一个错误,终止拷贝操作,err:", err)
		return
	}
	dst, err := os.OpenFile("../back.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("创建备份文件出现一个错误，终止拷贝操作,err:", err)
	}
	w, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("执行拷贝出现一个错误,此次拷贝操作终止,err:", err)
	}
	fmt.Println("此次备份的文件字节数为:", w, "字节")
}

func main() {
	copyFile()
}
