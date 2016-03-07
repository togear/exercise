package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type data struct { a int}
	var d = data{1234}
	var p *data

	p = &d

	fmt.Printf("%p, %v\n",p,p.a)

	x := 0x12345678

	q := unsafe.Pointer(&x)
	n := (*[4]byte)(q)

	for _,i := range n {
		fmt.Printf("%X",i)
	}

	for i := 0; i < len(n); i++ {
		fmt.Printf("%X",n[i])
	}
}
