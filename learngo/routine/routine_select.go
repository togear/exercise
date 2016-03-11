package main

import "fmt"

func main() {

	c := make(chan int)
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	for {
		select {
		case c <- 0: //没有语句，没有fallthrough
		case c <- 1:
		}
	}
}
