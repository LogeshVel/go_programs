package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("a: %v\n", a)
	// backup a to aa
	aa := make([]int, len(a))
	copy(aa, a)
	fmt.Printf("aa: %v\n", aa)
	fmt.Print("\nChanging value of a\n\n")
	a[0] = 100 // only a changes since the backing array of a and aa is different
	fmt.Printf("a: %v\n", a)
	fmt.Printf("aa: %v\n", aa)

	fmt.Print("\nAnother way to copy the slice\n\n")
	b := []int{100, 200, 300, 400, 500}
	fmt.Printf("b: %v\n", b)
	bb := append([]int(nil), b...)
	fmt.Printf("bb: %v\n", bb)
	fmt.Print("\nChanging value of b\n\n")
	b[0] = -1
	fmt.Printf("b: %v\n", b)
	fmt.Printf("bb: %v\n", bb)
}
