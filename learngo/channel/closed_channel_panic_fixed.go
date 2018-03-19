package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, "Send Result")
			case <-done:
				fmt.Println(idx, "Existing")
			}
		}(i)
	}

	fmt.Println("Result", <-ch)
	time.Sleep(1 * time.Second)
	close(done)
	time.Sleep(2 * time.Second)
}
