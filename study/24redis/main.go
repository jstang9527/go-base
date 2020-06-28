package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client
var rdb *redis.Cmdable

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}
	// fmt.Println("连接redis成功...")
	return nil
}

func setOrGet() {
	err := redisdb.Set("socre", 100, 0).Err()
	if err != nil {
		fmt.Println("set score failed, err:", err)
		return
	}
	v, err := redisdb.Get("socre").Result()
	if err != nil {
		fmt.Println("get socre failed, err:", err)
		return
	}
	fmt.Println("get score value: ", v)

	v2, err := redisdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
		// return
	} else if err != nil {
		fmt.Println("get name failed, err:", err)
		return
	} else {
		fmt.Println("name:", v2)
	}

	time.Sleep(time.Second * 3)
	v3, err := redisdb.Get("socre").Result()
	if err != nil {
		fmt.Println("get socre failed, err:", err)
		return
	}
	fmt.Println("get score value: ", v3)
}

func zset() {

	// fmt.Printf("zadd %d success.", num)

	// newScore, err := redisdb.ZIncrBy(key, 10.0, "Golang")
	// if err != nil {
	// 	return
	// }
	// fmt.Printf("golang's socre is %f now.\n", newScore)
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("连接redis成功...")

	// setOrGet()
	// key := "rank"
	// items := []*redis.Z{
	// 	&redis.Z{Score: 100.0, Member: "钟南山"},
	// 	&redis.Z{Score: 80.0, Member: "林医生"},
	// 	&redis.Z{Score: 70.0, Member: "王医生"},
	// 	&redis.Z{Score: 75.0, Member: "张医生"},
	// 	&redis.Z{Score: 59.0, Member: "叶医生"},
	// }
	// redisdb.ZAdd(key, &items...).Result()
}
