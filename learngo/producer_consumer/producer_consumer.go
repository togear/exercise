package main

import (
	"fmt"
)

func producer(c chan int, done chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Alice puts products, ID is %d \n", i)
		c <- i
	}

	done <- true

	defer close(c)
}

func consumer(c chan int) {
	hasMore := true
	var p int
	for hasMore {
		if p, hasMore = <-c; hasMore {
			fmt.Printf("Bob gets product,ID is %d\n", p)
		}
	}
}

func main() {
	done := make(chan bool)
	c := make(chan int)
	go producer(c, done)
	go consumer(c)
	<-done
}
