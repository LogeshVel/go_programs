package main

import "fmt"

func main() {
	var a int = 100

	p := &a
	fmt.Printf("p: %v\n", p)
	fmt.Printf("*p: %v\n", *p)

	fmt.Printf("a: %v\n", a)
	aa := *p
	fmt.Printf("aa: %v\n", aa)
	aa = 1111
	fmt.Printf("aa: %v\n", aa)

	fmt.Printf("*p: %v\n", *p)
	fmt.Printf("a: %v\n", a)
	*p = 999
	fmt.Printf("*p: %v\n", *p)
	fmt.Printf("a: %v\n", a)

}
