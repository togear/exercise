package main

import "fmt"

type T struct {
	Foo string
	Bar int
	Qux string
}

func main() {
	t := T{Foo: "example", Bar: 123}
	fmt.Printf("t %+v\n", t)
}
