package main

import "fmt"

func main() {
	var s []int
	printSlice("s",s)

	s = append(s,0)
	printSlice("s",s)

	s = append(s,1)
	printSlice("s",s)

	s = append(s,2,3,4)
	printSlice("s",s)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len =%d cap= %d %v\n",
				s , len(x),cap(x),x)
}
