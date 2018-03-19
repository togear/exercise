package main

import "fmt"

func main() {
	x := "text"
	//	x[0] = 'T'
	//	fmt.Println(x)
	// ./string_modify.go:7: cannot assign to x[0]
	//	xBytes := []byte(x)
	//	xBytes[0] = 'T'
	//	x = string(xBytes)
	//	fmt.Println(x)

	xRunes := []rune(x)
	xRunes[0] = '我'
	x = string(xRunes)
	fmt.Println(x)
}
