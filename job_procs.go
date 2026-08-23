package main

import (
	"fmt"
	"sync"
)

func worker(i int, jobs chan Job, results chan Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker %d initiated\n", i)
		fmt.Printf("worker %d recieved job %v\n", i, job.ID)

		result := Result{
			job:    job,
			output: "job finished",
		}

		fmt.Printf("worker %d procesing job %v\n", i, job.ID)

		results <- result
		fmt.Printf("worker %d pushed job into result channel. result: %v\n", i, result)
	}

}
