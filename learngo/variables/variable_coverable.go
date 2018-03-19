package main

// go tool vet -shadow variable_coverable.go
//variable_coverable.go:8: declaration of "x" shadows declaration at variable_coverable.go:4

func main() {
	x := 1
	println(x)
	{
		println(x)
		x := 2
		println(x)
	}

	println(x)
}
