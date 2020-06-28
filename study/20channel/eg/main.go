package main

import (
	"fmt"
	"sync"
)

var c chan int //引用类型，需要初始化，否则为nil 需要指定通道的元素类型
var wg sync.WaitGroup

func main() {
	fmt.Println(c)     //nil
	c = make(chan int) //无缓冲通道初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-c)
	}()
	c <- 99
	wg.Wait()

	c = make(chan int, 1)
	c <- 100
	fmt.Println(<-c)
	close(c)
}
