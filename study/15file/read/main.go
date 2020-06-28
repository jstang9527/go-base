package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFileByBytes() {
	fileObj, err := os.Open("./main.go")
	defer fileObj.Close()
	if err != nil {
		fmt.Println("打开文件出错:", err)
		return
	}
	var tmpBytes [128]byte
	for {
		n, err := fileObj.Read(tmpBytes[:]) //n读取到的字节
		if err == io.EOF {
			fmt.Println("文件数据已读取完毕:", err)
			return
		}
		if err != nil {
			fmt.Println("read from file failed, err:", err)
			return
		}
		fmt.Println("读取到", n, "字节")
		fmt.Println(string(tmpBytes[:n]))
		if n < 128 { //本次没读满,意味着已经没数据了,可以退出
			return
		}
	}
}

//一行行读
//bufio是在file的基础上封装了一层API,支持更多的功能
func readFileByBufio() {
	fileObj, err := os.Open("./main.go")
	defer fileObj.Close()
	if err != nil {
		fmt.Println("文件打开错误, err: ", err)
		return
	}
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件已读完, 标志:", err)
			return
		}
		if err != nil {
			fmt.Println("读取文件发生一个错误, err: ", err)
			return
		}
		fmt.Print(line)
	}
}

//读取整个文件
func readFileByIoutil() {
	content, err := ioutil.ReadFile("main.go")
	if err != nil {
		fmt.Println("读取文件发生错误, err:", err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	readFileByIoutil()
}
