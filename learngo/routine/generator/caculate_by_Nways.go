package main

import (
	"fmt"
	"math/rand"
	"time"
)

func do_stuff(x int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	return 100 - x
}

func branch(x int) chan int {
	ch := make(chan int)
	go func() {
		ch <- do_stuff(x)
	}()
	return ch
}

func fanIn(branches ...chan int) chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < len(branches); i++ {
			select {
			case v1 := <-branches[i]:
				ch <- v1
			}
		}
	}()

	return ch
}

func main() {
	result := fanIn(branch(1), branch(2), branch(3), branch(4))

	for i := 0; i < 4; i++ {
		fmt.Println(<-result)
	}
}
