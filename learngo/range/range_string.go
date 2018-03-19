package main

import "fmt"

func main() {
	data := "A\xfe\x02\xff\x04"
	for _, v := range data {
		fmt.Printf("%#x ", v) // 0x41 0xfffd 0x2 0xfffd 0x4	// 错误
	}

	fmt.Println("")

	for _, v := range []byte(data) {
		fmt.Printf("%#x ", v) // 0x41 0xfe 0x2 0xff 0x4	// 正确
	}
	fmt.Println("")
}
