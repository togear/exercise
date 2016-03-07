package main

func test(x, y int) {
	var z int

	func() {
		defer func() {
			if recover() != nil { z = 0 }
		}()

		z = x / y
		return
	}()

	println("x / y =", z)
}

func main() {
	test(10,2)
	test(10,0)
}

