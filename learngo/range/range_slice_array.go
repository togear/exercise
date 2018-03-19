// range
package main

import (
	"fmt"
)

func main() {
	x := []string{"a", "b", "c"}

	//	for  v := range x { //do not make mistake
	for _, v := range x {
		fmt.Println(v)
	}
}
