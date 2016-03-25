package main

import "fmt"

//https://tour.golang.org/basics/14

func main() {
	v := 42  //change me
	fmt.Printf("v is of type %T\n",v)

	f := 43.2
	fmt.Printf("v is of type %T\n",f)

	v = 'c'
	fmt.Printf("v is of type %T\n",v)
}
