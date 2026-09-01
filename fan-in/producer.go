package fanin

import (
	"math/rand"
	"sort"
	"time"
)

func producer(id, count int) (ichan chan int) {
	out := make(chan int)
	go func() {
		defer close(out)
		vals := make([]int, count)
		for i := range vals {
			vals[i] = rand.Intn(100)
		}
		sort.Ints(vals) // each producer's own stream must be sorted
		for _, v := range vals {
			time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
			out <- v
		}
	}()
	return out
}
