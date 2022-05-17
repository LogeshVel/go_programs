package main

import "fmt"

func main() {
	// var a1 = [5]int{5, 3, 9, 2, 6}
	var s1 = []int{} // initialized but empty
	var s2 []int     // uninitilaized so nil
	fmt.Printf("s1: %#v\n", s1)
	fmt.Printf("len(s1): %d\n", len(s1))
	fmt.Printf("s2: %#v\n", s2)
	fmt.Printf("len(s2): %d\n", len(s2))
	// fmt.Printf("By Index - a1[0]  : %d   and the Type is   %T\n", a1[0], a1[0])
	// fmt.Printf("By Slice - a1[0:1]: %d and the Type is %T\n", a1[0:1], a1[0:1])

	// fmt.Printf("a1: %#v\n", a1)
	// fmt.Printf("s1: %#v\n", s1)

	// sort.Ints(a1[:])
	// sort.Ints(s1)
	// fmt.Print("\nAfter sorting\n\n")
	// fmt.Printf("a1: %#v\n", a1)
	// fmt.Printf("s1: %#v\n", s1)

}
