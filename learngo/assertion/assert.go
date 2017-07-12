package main

import (
	"fmt"
)

func Test(a interface{}) string {
	value, ok := a.(string)
	if !ok {
		fmt.Println("It's not ok for type string", value)
		return ""
	}
	return value
}

/*
func Test_error(a interface{}) string {
	return string(a)
}
*/

func main() {
	a := "abcd"
	c1 := Test(a)
	b := 123
	c2 := Test(b)
	fmt.Println(c1, c2)

	/*
		a = "abcd"
		c1 = Test_error(a)
		b = 123
		c2 = Test_error(b)
		fmt.Println(c1, c2)
	*/
}
