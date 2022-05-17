package main

import (
	"fmt"
)

func main() {

	// key := "something"
	a := map[string]int{}
	a["zero"] = 0
	a["one"] = 1

	// how to check whether the key exists or not
	// value, ok := a[key]
	// if ok {
	// 	fmt.Printf("The key %q exists\n", key)
	// 	fmt.Printf("a[%q] : %v\n", key, value)
	// } else {
	// 	fmt.Printf("The key %q not exists\n", key)
	// 	fmt.Printf("The zero value of a[%q] : %v\n", key, value)
	// }

	// to compare 2 maps
	b := map[string]int{"one": 1, "zero": 0}

	aStr := fmt.Sprintf("%v", a)
	bStr := fmt.Sprintf("%v", b)
	if aStr == bStr {
		fmt.Println("both are equal")
		fmt.Println(a)
		fmt.Println(b)
	} else {
		fmt.Println("both are not equal")
		fmt.Println(a)
		fmt.Println(b)
	}

}
