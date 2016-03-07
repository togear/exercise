package main

import (
	"fmt"
	"unsafe"
)

func main() {
	d := struct {
		s string
		x int
	}{"abc",100}

	p := uintptr(unsafe.Pointer(&d))
	p += unsafe.Offsetof(d.x)

	p2 := unsafe.Pointer(p)
	px := (*int)(p2)
	*px = 200

	fmt.Printf("%#v\n",d)
}
