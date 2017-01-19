package main

import (
	"fmt"
)

func main() {
	c, quit := make(chan int), make(chan int)

	go func() {
		c <- 2
		quit <- 1
	}()

	for is_quit := false; !is_quit; {
		select {
		case v := <-c:
			fmt.Printf("received %d from c\n", v)
		case <-quit:
			is_quit = true
		}
	}

}
