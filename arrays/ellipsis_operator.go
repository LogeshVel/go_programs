package main

import "fmt"

func main() {
	// In an array , the ... notation specifies a length equal to the number of elements
	// initialized in the array literal.
	a := [...]int{1, 2, 3, 4, 5, 6}
	a[5] = -6 // this is fine becoz the 5 is valid index for the 6 elements array
	fmt.Printf("%v\n", a)
	fmt.Printf("%#v\n", a)
	// a[6] = -7 this is error becoz list has only 6 elements(0 to 5)
	// append function won't wrk in array
	// a = append(a, -100)

	// keyed array
	fmt.Println("working with keyed arrays")
	k := [...]string{
		9: "10th ele",
	}
	// here 9 is the index of that value and the values for other indexes
	// before 9 were zero valued
	fmt.Printf("k: %#v\n", k)

	kk := [5]int{
		3: 4,
		5,
		1: 2,
	}
	// here the unkeyed value 5 will take the index from the last keyed element.
	// so 5 will be in the index of 4(last)
	fmt.Printf("kk: %#v\n", kk)
}
