package main

import (
	"fmt"
	"strconv"
	"sync"
)

// var m = make(map[string]int, 10)
var m = sync.Map{}

// func get(key string) int {
// 	return m[key]
// }
// func set(key string, value int) {
// 	m[key] = value
// }

func main() {
	wg := sync.WaitGroup{}
	// for i := 0; i < 49; i++ {
	// 	wg.Add(1)
	// 	go func(n int) {
	// 		key := strconv.Itoa(n)
	// 		set(key, n)
	// 		fmt.Printf("k=:%v,v=:%v\n", key, get(key))
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n) //存的时候会先看map有无锁
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v=:%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
