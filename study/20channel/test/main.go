package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//读取任务，获取结果
func sendNum(id int, jobs <-chan int, results chan<- int) {
	rand.Seed(time.Now().UnixNano())
	for j := range jobs { //从jobs管道获取任意指令
		num := rand.Intn(10)
		fmt.Printf("worker:%d start job:%d send[%d]\n", id, j, num)
		time.Sleep(time.Second * time.Duration(num))
		results <- num
	}
	wg.Done()
}

var jobs chan int
var results chan int
var wg sync.WaitGroup

func main() {
	jobs = make(chan int, 100)
	results = make(chan int, 100)
	go func() { //有多少取多少，直到result关闭
		for a := range results { //如果results没关闭，会一直阻塞，引发deadlock panic
			fmt.Println("接收成果：", a)
		}
	}()
	//开启10个goroutine
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go sendNum(w, jobs, results)
	}
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)    //任务发放完毕可以关闭发生通道
	wg.Wait()      //等待工人完成工作
	close(results) //工人已全部完成工作，关闭接收通道

}
