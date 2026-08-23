package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerPool(wg *sync.WaitGroup) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel()

	jobs := make(chan Job)
	results := make(chan Result)

	for i := range 10 {
		wg.Add(1)
		go worker(ctx, i, jobs, results, wg)
	}

	go producer(jobs, 20)

	go func() {
		wg.Wait()
		fmt.Println("done waiting")
		close(results)

	}()

	finished := []int{}
	for r := range results {
		fmt.Println(r)
		finished = append(finished, r.job.ID)
		fmt.Printf("result[%d]: %v\n", r.job.ID, r.output)
	}
	fmt.Println(finished)
}
