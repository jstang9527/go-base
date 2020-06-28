package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func f2(ctx context.Context) {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("yyyy")
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
FORLOOP:
	for {
		fmt.Println("xxxx")
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)

	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()

	//如何通知子goroutine退出？
}
