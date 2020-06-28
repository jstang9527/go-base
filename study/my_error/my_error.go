package main

import "fmt"

func openDB()  { fmt.Println("连接数据库") }
func closeDB() { fmt.Println("关闭数据库") }

func openFile() {
	defer func() { //defer一定要在可能引发的panic之前定义
		if err := recover(); err != nil { //recover()必须搭配defer使用
			fmt.Println(err)
		}
		closeDB()
	}()
	panic("没有该文件")
	fmt.Println("sadad") //出异常则不可到达的代码
}

func main() {
	// openDB()
	// openFile()
	var (
		name  string
		age   int
		class string
	)
	// fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)
}
