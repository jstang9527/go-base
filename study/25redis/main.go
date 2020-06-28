package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var conn redis.Conn

func initRedis() {
	var err error
	conn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect to redis failed, err: ", err)
		return
	}
	fmt.Println("connect redis server success.")
	// defer conn.Close()
}

func zset() {
	//insert zset
	_, err := conn.Do("ZADD", "mykey", "INCR", 1, "robot1")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	// 再执行一个有序zset插入
	_, err = conn.Do("ZADD", "mykey", "INCR", 1, "robot2")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	// 读取指定zset
	userMap, err := redis.StringMap(conn.Do("ZRANGE", "mykey", 0, 10, "withscores"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", userMap)
	}

	for user := range userMap {
		fmt.Printf("user name: %v %v\n", user, userMap[user])
	}

}

func main() {
	initRedis()
	zset()
}
