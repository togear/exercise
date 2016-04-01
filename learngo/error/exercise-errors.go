package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt)Error() string {
//	return fmt.Sprintf("cannot Sqrt negative number: %f",float64(e))
	return fmt.Sprintf("cannot Sqrt negative number: %f",e)
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return math.Sqrt(x),nil
	}	
	return x, ErrNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
