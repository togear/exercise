package main

import(
	"fmt"
	"math/rand"
)

func main(){
	r := rand.New(rand.NewSource(99))
	fmt.Println(r)

}
