package main

import "fmt"

//面向对象把方法和属性数据放到一个类里面
//面向接口,不关心数据是什么,变量只要实现了接口的方法=变量实现了接口,即该变量为接口类型的变量

type mysql struct {
	host     string
	ip       string
	port     uint32
	user     string
	password string
}
type oracle struct {
	host     string
	ip       string
	port     uint32
	user     string
	password string
}

func (m mysql) connect() {
	fmt.Printf("连接%v数据库\n", m.host)
}
func (o oracle) connect() {
	fmt.Printf("连接%v数据库\n", o.host)
}

type db interface {
	connect()
}

//连接数据库函数
func connectDB(d db) {
	d.connect()
}

func main() {
	// o1 := oracle{host: "oracal"}
	var o1 oracle
	m1 := mysql{host: "mysql"}
	connectDB(o1)
	connectDB(m1)
}
