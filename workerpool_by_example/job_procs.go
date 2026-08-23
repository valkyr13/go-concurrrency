package workerpool

import (
	"context"
	"fmt"
	"sync"
)

func worker(ctx context.Context, i int, jobs chan job, results chan result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d initiated\n", i)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("error: %v\n", ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("worker %d recieved channel close signal. No more jobs are coming\n", job.ID)
				return
			}
			fmt.Printf("worker %d recieved job %v\n", i, job.ID)

			result := result{
				job:    job,
				output: "job finished",
			}
			fmt.Printf("worker %d procesing job %v\n", i, job.ID)

			results <- result
			fmt.Printf("worker %d pushed job into result channel. result: %v\n", i, result)

		}

	}

}
