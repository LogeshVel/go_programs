package main

import "fmt"

func main() {
	var trunced_float float64
	fmt.Printf("Enter a float number: ")
	fmt.Scan(&trunced_float)
	fmt.Printf("Truncated version of the given float value : %.f\n", trunced_float)
}
