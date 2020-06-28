package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeFile() {
	fileObj, err := os.OpenFile("../log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666) //没有则新建、追加、只写
	defer fileObj.Close()
	if err != nil {
		fmt.Println("读取文件发生一个错误", err)
	}
	if n, err := fileObj.Write([]byte("open the file2\n")); err != nil {
		fmt.Println("出错了,err:", err)
	} else {
		fmt.Printf("写入%d字节\n", n)
	}
	if n, err := fileObj.WriteString("close file2"); err != nil {
		fmt.Println("出错了, err:", err)
	} else {
		fmt.Printf("写入%d字节\n", n)
	}
}

func writeFileByBufio() {
	fileObj, err := os.OpenFile("../log.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer fileObj.Close()
	if err != nil {
		fmt.Println("读取文件发生一个错误,err:", err)
		return
	}
	write := bufio.NewWriter(fileObj)
	write.WriteString("how are you!") //写入缓冲
	write.Flush()                     //写入文件
}

func writeFileByIoutil() {
	str := "i am funny."
	if err := ioutil.WriteFile("../log.txt", []byte(str), 0666); err != nil {
		fmt.Println("写入文件发生个错误, err:", err)
		return
	}
}

func main() {
	// writeFileByBufio()
	writeFileByIoutil()
}
