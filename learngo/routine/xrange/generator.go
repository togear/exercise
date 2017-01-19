package main

import (
	"fmt"
)

func xrange() chan int {
	var ch chan int = make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func main() {
	generaror := xrange()

	for i := 0; i < 1000; i++ {
		fmt.Println(<-generaror)
	}
}
