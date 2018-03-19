package main

import "fmt"

func main() {
	h, w := 2, 4
	raw := make([]int, h*w)

	for i := range raw {
		raw[i] = i
	}

	fmt.Println(raw, &raw[4])

	table := make([][]int, h)
	for i := range table {
		//0: raw[0*4: 0*4 + 4]
		//1: raw[1*4: 1*4 + 4]
		table[i] = raw[i*w : i*w+w]
	}

	fmt.Println(table, &table[1][0])
}
