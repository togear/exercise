package main

func main() {
	recover()
	panic("not good")
	recover()
	println("ok")
}
