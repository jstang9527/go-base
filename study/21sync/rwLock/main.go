package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0

// var lock sync.Mutex
var wg sync.WaitGroup
var rwLock sync.RWMutex

func read() {
	defer wg.Done()
	rwLock.RLock()
	fmt.Print(x)
	time.Sleep(time.Millisecond) //读花费1毫秒
	rwLock.RUnlock()

}

func write() {
	defer wg.Done()
	rwLock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5) //写花费5毫秒
	rwLock.Unlock()
}

func main() {
	start := time.Now()
	wg.Add(1010)
	for i := 0; i < 10; i++ {
		go write()
	}
	for i := 0; i < 1000; i++ {
		go read()
	}
	wg.Wait()
	fmt.Print("\n", x)
	fmt.Println(time.Now().Sub(start))

}
