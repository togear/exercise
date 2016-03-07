package main

import "fmt"

var x,y,z int;
var s,n = "abc", 123

var (
	a int
	b float32
)

func test()(int,string) {
	return 1,"abc"
}

type Color int
const (
	Black Color = iota
	Red
	Blue

)

func test_color(c Color){
	fmt.Println(c)
}

func main() {
	n,s := 0x1234, "Hello World!"
	fmt.Println(x,s,n)

	_, s1 := test()
	fmt.Println(s1)

	c := Black
	test_color(c)

	x := 1
//	test_color(x)  //cannot use x (type int) as type Color in argument to test_color
	_ = x
	test_color(1)
}
