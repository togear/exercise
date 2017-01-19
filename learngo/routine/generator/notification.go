package main

import (
	"fmt"
)

func get_notification(user string) chan string {
	notifications := make(chan string)

	go func() {
		notifications <- fmt.Sprintf("Hi %s,wlecome to weibo.com!", user)
	}()

	return notifications
}

func main() {
	jack := get_notification("jack")
	joe := get_notification("joe")

	fmt.Println(<-jack)
	fmt.Println(<-joe)
}
