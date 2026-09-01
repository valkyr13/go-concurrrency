package fanin

import (
	"container/heap"
	"sync"
)

func KWayMerge(ichans []chan int, ochan chan int) {
	// n go routine - recieving from n channels

	n := len(ichans)
	var wg sync.WaitGroup
	h := &MinHeap{}
	heap.Init(h)

	var mu sync.Mutex

	for i := 0; i < n; i++ {
		ichans[i] = producer(i, 5) // 4 producers, 5 sorted ints each
	}

	for i := range n {
		wg.Add(1)
		go consumer(&wg, ichans[i], h, &mu)
	}

	wg.Wait()

	// wait until all input channels are done pushing and closed
	j := h.Len()

	for ; j > 0; j-- {
		el := heap.Pop(h)
		ochan <- el.(int)
	}
	close(ochan)

}
