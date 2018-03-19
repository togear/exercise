package main

import "fmt"

// 错误的调用示例
func main() {
	defer func() {
		doRecover()
	}()
	panic("not good")
}

func doRecover() {
	fmt.Println("recobered: ", recover())
}
