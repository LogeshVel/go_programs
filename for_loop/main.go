package main

import "fmt"

func main() {
	a := 100
	// no initialize and post statement but the condition is hidden (default) true
	// its an infinite loop
	for {
		fmt.Println("In Loop")
		if a == 103 {
			fmt.Println("Breaking the Loop")
			break
		}
		a++
	}
}
