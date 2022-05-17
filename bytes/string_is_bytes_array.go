package main

import "fmt"

func main() {
	var byteSlice []byte
	var stringValue string = "Logesh"
	// when I try to append the string to the bytes slice
	// its an error
	// but String is byte slice behind the scene
	byteSlice = append(byteSlice, stringValue...)
	fmt.Printf("byteSlice        : %v\n", byteSlice)
	fmt.Printf("string(byteSlice): %v\n", string(byteSlice))
	fmt.Printf("byteSlice by %%s  : %s\n", byteSlice)
}
