// array
package main

import (
	"fmt"
)

func main() {
	a := [] int {1,2,3,4,8,9}
	for i := range a {
		fmt.Println( i,"----",a[i])
	}
}
