package main

import "fmt"

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	stream := pump()
	fmt.Println(<-stream)
}
