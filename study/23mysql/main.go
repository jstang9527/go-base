package main

import (
	//只有方法标签，没有具体实现，具体实现交给其他数据库驱动(mysql/oracle)
	"database/sql" //会自动创建和释放连接；它也会维护一个闲置连接的连接池。
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//sql.DB是一个数据库（操作）句柄，代表一个具有零到多个底层连接的连接池。它可以安全地被多个goroutine同时使用。
var db *sql.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return (err)
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(5)
	fmt.Println("连接数据库成功！！！")
	return nil
}

type user struct {
	id   int
	name string
	age  int
}

//查单条记录
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.QueryRow(sqlStr, 3).Scan(&u.id, &u.name, &u.age) //从连接池拿连接
	if err != nil {                                            //没有该记录
		fmt.Println("scan failed, err: ", err)
		return
	}
	fmt.Printf("id:%v name:%v age:%v\n", u.id, u.name, u.age)
	for i := 0; i < 10; i++ {
		fmt.Println("模拟拿连接不释放", i)
		db.QueryRow(sqlStr, 1) //第6次阻塞
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err: %v\n", err)
		return
	}
	//1.查询单条记录sql
	queryRowDemo()
}
