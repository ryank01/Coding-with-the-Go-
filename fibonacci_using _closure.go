package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// this demonstrates the use of a closure
func fibonacci() func(int) int {
	first := 0
	second := 1
	return func(x int) int {
		next := first + second
		first = second
		second = next
		return	next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
