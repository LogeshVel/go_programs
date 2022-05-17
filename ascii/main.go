package main

import (
	"fmt"
)

func main() {
	fmt.Println("Get the ASCII code of your character. Type quit to quik the Program")
	fmt.Println("Enter a character")
	var char string
	for {
		fmt.Scanln(&char)
		if char == "quit" {
			fmt.Println("Quitting the program")
			return
		}
		fmt.Printf("ASCII code for %c is %d", char[0], rune(char[0]))
	}

}
