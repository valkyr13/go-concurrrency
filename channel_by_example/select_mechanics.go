package channel

import (
	"context"
	"fmt"
	"time"
)

func RandomSelectPick() {

	for i := 0; i < 5; i++ {
		//fresh initialisation, drains the channels automatically
		ch1, ch2 := make(chan string, 1), make(chan string, 1)

		ch1 <- "incoming channel 1" // will be blocked here until someone recieves it
		ch2 <- "incoming channel 2"

		select {
		case v := <-ch1:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		}
	}

}

func NothingReadyWithDefault() {
	ch := make(chan int)
	go func() {
		fmt.Println(<-ch) //blocks
	}()
	ch <- 1 //blocks

}

func TimeoutPattern(ch chan int, ctx context.Context) {
	select {
	case v := <-ch:
		fmt.Println("got :", v)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timed out")
	case <-ctx.Done():
		fmt.Println("cancelled: ", ctx.Err())
	}

}
