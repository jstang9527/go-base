package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://192.168.231.128:10001/xxx/?user=go&password=12345")
	if err != nil {
		fmt.Println("连接客户端发生错误, err: ", err)
		return
	}
	defer resp.Body.Close()
	bData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取数据出错, err:", err)
		return
	}
	fmt.Println(string(bData))
}
