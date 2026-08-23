package main

import "fmt"

func producer(jobs chan Job, n int) {
	for i := range n {
		fmt.Printf("pushing job %d into job channel\n", i)
		jobs <- Job{
			ID: i,
		}
	}
	close(jobs)
	// Note : With an unbuffered channel, the producer can only send as many jobs simultaneously as there are workers actively waiting to receive.

}
