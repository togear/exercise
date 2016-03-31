package main
import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first := 0
	second := 0
	last_1 := 0
	last_2 := 0
	now := 0
	return func() int {
		if first == 0 {
			first ++;
			return last_1;
			
		}
		if second == 0 {
			last_2 = first + second
			second++;
			return last_2;
		}

		now = last_1 + last_2;
		last_1 = last_2
		last_2 = now

		return now 
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
