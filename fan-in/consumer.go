package fanin

import (
	"container/heap"
	"sync"
)

func consumer(wg *sync.WaitGroup, ichan chan int, h *MinHeap, mu *sync.Mutex) {
	defer wg.Done()

	for i := range ichan {
		mu.Lock()
		heap.Push(h, i)
		mu.Unlock()
	}

	// close(ichan) sender shouldn't close chan
}
