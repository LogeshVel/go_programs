package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// fmt.Println(s)
	// fmt.Printf("s: %#v\n", s)
	// println(s == nil)
	a := s[0:3]
	// fmt.Println(a == s) // invalid operation: a == s (slice can only be compared to nil)
	res := append(s, a...)
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(res)

}
