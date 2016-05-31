package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Testing mergesort...")
	data := gen_data(1000)
	fmt.Println("before:", data)
	sorted := mergesort(data)
	fmt.Println("after:", sorted)
	err := false
	for i, val := range sorted {
		if i != val {
			err = true
			fmt.Println("error:", i, "does not equal", val)
		}
	}
	if !err {
		fmt.Println("pass")
	}
	fmt.Println("Done testing mergesort...")
}

func gen_data(size int) []int {
	return rand.Perm(size)
}

func mergesort(arr []int) (out []int) {
	out = make([]int, len(arr))
	sorted := make(chan int)
	go mergesort2(sorted, arr)
	for i := range out {
		out[i] = <-sorted
	}
	return
}

func mergesort2(out chan int, arr []int) {
	if len(arr) <= 1 {
		out <- arr[0]
		out <- math.MaxInt32 //reportMergeComplete
		return
	}
	a := make(chan int)
	b := make(chan int)
	go mergesort2(a, arr[:len(arr)/2])
	go mergesort2(b, arr[len(arr)/2:])
	v1 := <-a
	v2 := <-b
	for _ = range arr {
		if v1 < v2 {
			out <- v1
			v1 = <-a
		} else {
			out <- v2
			v2 = <-b
		}
	}
	out <- math.MaxInt32 //reportMergeComplete
}
