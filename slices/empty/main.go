package main

import "fmt"

func main() {
	// uninitialized slice is nil
	var s1 []string
	// initialized slice ith zero element
	var s2 = []string{}

	fmt.Printf("s1: %#v\n", s1)
	fmt.Printf("s2: %#v\n", s2)

	if s1 == nil {
		fmt.Println("s1 is nil")
	} else {
		fmt.Println("s1 is not nil")
	}

	if s2 == nil {
		fmt.Println("s2 is nil")
	} else {
		fmt.Println("s2 is not nil")
	}
	// An Empty slice is an initialized slice
	// An Nil slice is an un-initialized slice
	fmt.Println("len of s1:", len(s1))
	fmt.Println("len of s2:", len(s2))
	//Both the slice has same length - 0
}
