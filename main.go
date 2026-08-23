package main

import (
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
		done <- grow(200000)
	}()

	fmt.Println(<-done)

	// worker pool pratice

	workerPool(&wg)

}
