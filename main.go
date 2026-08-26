package main

import (
	channel "concurrency/channel_by_example"
	workerpool "concurrency/workerpool_by_example"
	"context"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//wait group practice

	runtime.GOMAXPROCS(runtime.NumGoroutine())
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("goroutine", id, "alive; total:", runtime.NumGoroutine())
		}(i)
	}
	wg.Wait()

	done := make(chan int)

	go func() {
		done <- channel.Grow(200000)
	}()

	fmt.Println(<-done)

	// worker pool pratice
	workerpool.WorkerPool(&wg)

	// select mechanics
	channel.RandomSelectPick()
	channel.NothingReadyWithDefault()
	ctx := context.Background()
	ch := make(chan int)

	go channel.TimeoutPattern(ch, ctx)
	ch <- 1

}
