package main

import "fmt"

func main() {
	a := 2
	fmt.Println("a is", a)
come_here:
	if a <= 5 {
		fmt.Println("a is(", a, ") less than 5 so incrementing it")
		a++
		goto come_here
	} else {
		fmt.Println("a(", a, ") is greater than 5")
		goto go_here
	}

	// fmt.Println("outer")

go_here:
	fmt.Println("Finally in go here")

}
