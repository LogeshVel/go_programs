package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	start := "i"
	end := "n"
	fmt.Printf("Enter a string: ")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	line = strings.TrimSpace(line)
	var lowerS string = strings.ToLower(string(line))
	if string(lowerS[0]) == start && string(lowerS[len(lowerS)-1]) == end && strings.Contains(lowerS, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")

	}

}
