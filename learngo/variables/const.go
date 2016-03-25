//const.go
package main

import (
	"fmt"
)
const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

const Pi = 3.14


func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	const World = "世界"
	fmt.Println("Hello",World)
	fmt.Println("Happy",Pi,"Day")

	const Truth = true
	fmt.Println("Go rules?",Truth)


}
