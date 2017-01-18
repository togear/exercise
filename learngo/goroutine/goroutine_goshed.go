// main.go
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 2; i++ {

		if s == "hello" {

			fmt.Println("~~ hello")
		} else {

			fmt.Println("~~ world")
		}

		runtime.Gosched()
		fmt.Println(s)

		if s == "hello" {

			fmt.Println("2~~ hello")
		} else {

			fmt.Println("2~~ world")
		}

	}
}

func main() {
	runtime.GOMAXPROCS(1)

	go say("world")

	say("hello")

}
