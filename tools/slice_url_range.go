package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Slice Range")

	var fileSize int64 = 1024*1024 + 2
	var sliceSize int64 = 1024 * 1024
	numf := float64(fileSize) / float64(sliceSize)
	num := int64(math.Ceil(float64(numf)))
	fmt.Println("at range", numf)
	fmt.Println("at range", num)

	url := "http://www.test.com/index.html"
	rangeHeader := "bytes="

	fmt.Println("slice file size", fileSize, "for", sliceSize, "At Num", num)

	for i := int64(0); i < num; i++ {
		fmt.Println("slice file size", fileSize, "for", sliceSize, "At Num [", num, "] ", i)
		start := i * sliceSize
		end := (i+1)*sliceSize - 1
		var rangeValue string = fmt.Sprintf("%d-%d", start, end)
		fmt.Println("at range", url)
		fmt.Println("at range", rangeHeader)
		fmt.Println("at range", rangeValue)
	}

}
