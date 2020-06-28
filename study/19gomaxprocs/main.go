package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Print("A:", i)
		time.Sleep(time.Millisecond * 300)
	}
}
func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Print("B:", i)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}