package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func hello() {
	rand.Seed(time.Now().UnixNano()) //保证每次执行的时候都不一样
	for i := 0; i < 5; i++ {
		r1 := rand.Int() //int64随机数
		r2 := rand.Intn(100)
		fmt.Println(r1, r2)
	}

}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()
}
