package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Loop var i=", i)
		i := i + 2
		fmt.Println("Inside Loop var i=", i)
	}

	for x := 1; x == 10; x++ {
		fmt.Println("a")
	}
	// fmt.Println("outer:", x) // this is error
}
