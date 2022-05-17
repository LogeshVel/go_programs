package main

import "fmt"

func main() {
	a := 10.78
	fmt.Printf("a is : %f\n", a)
	fmt.Printf("a is : %e\n", a)
	fmt.Printf("a is : %g\n", a)
	// %g is used in place of %e and %f
	// where we need to have the
	// decimal values as given
	// no extra or no cutoff decimal values
}
