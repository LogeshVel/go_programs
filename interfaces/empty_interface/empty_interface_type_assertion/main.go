package main

import (
	"fmt"
	"strings"
)

func main() {

	var e_i interface{}
	fmt.Printf("Type of e_i: %T\n", e_i)
	e_i = "hello"
	fmt.Println(e_i)
	fmt.Printf("Type of e_i: %T\n", e_i)

	// fmt.Println(strings.Contains(e_i.(string), "he"))

	s, ok := e_i.(string)
	if ok {
		fmt.Println(strings.Contains(s, "he"))
	}
	e_i = 10
	fmt.Println(e_i)
	fmt.Printf("Type of e_i: %T\n", e_i)

	// s := e_i.(string)
	// fmt.Println(s)

	// s, ok := e_i.(string)
	// fmt.Println(s, ok)

	// f, ok := i.(float64)
	// fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}
