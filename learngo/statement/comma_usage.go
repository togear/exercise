package main

func main() {
	x := []int{
		1,
		2, // syntax error: unexpected newline, expecting comma or }
	}
	y := []int{1, 2}
	z := []int{1, 2}

	_ = x
	_ = y
	_ = z
	// ...
}
