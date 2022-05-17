package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "String \"String\""
	fmt.Println(s1)

	raw_string := `Raw String can be wriiten within "back ticks" \n \t won't wrk`
	fmt.Println(raw_string)

	s := "Hi there  Go!"

	fmt.Printf("%s\n", s) // => Hi there  Go!
	fmt.Printf("%q\n", s) // => "Hi there  Go!"

	fmt.Println("The First index has -", s[1])
	fmt.Printf("The First index has - %c\n", s[1])

	str := "ţară" // ţară means country in Romanian
	// 't', 'a' ,'r' and 'a' are runes and each rune occupies beetween 1 and 4 bytes.

	//The len() built-in function returns the no. of bytes not runes or chars.
	fmt.Println(len(str)) // -> 6,  4 runes in the string but the length is 6

	// returning the number of runes in the string
	m := utf8.RuneCountInString(str)
	fmt.Println(m) // => 4

}
