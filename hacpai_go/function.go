package main

import "fmt"

func test() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

func test_slice(s string, n ...int) string {
	var x int

	for _, i := range n {
		x += i
	}

	return fmt.Sprintf(s,x)
}

func main() {
	f := test()
	f()

	fmt.Println(test_slice("sum :‘%d",1,2,3,4))
}

