package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "ţară" // ţară means country in Romanian
	// 't', 'a' ,'r' and 'a' are runes and each rune occupies beetween 1 and 4 bytes.

	//The len() built-in function returns the no. of bytes not runes or chars.
	fmt.Println("len(str):", len(str)) // -> 6,  4 runes in the string but the length is 6

	// returning the number of runes in the string
	m := utf8.RuneCountInString(str)
	fmt.Println("RuneCountInString :", m) // => 4
	// wont work in the way we expect, becoz this loops over byte by byte not char by char
	fmt.Printf("\n Using for loop with len(str)\n\n")
	for n := 0; n < len(str); n++ {
		fmt.Println("Loop count:", n)
		fmt.Printf("str[%d]: %c\n", n, str[n])
	}

	// range loop works as expected
	fmt.Printf("\n Using range loop\n\n")
	for idx, c := range str {
		fmt.Printf("index : %d , %c\n", idx, c)
	}

}
