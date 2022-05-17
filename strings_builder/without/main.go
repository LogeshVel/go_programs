package main

import "fmt"

func main() {
	var s = []string{"One", "Two", "Three", "Four", "Five"}
	stringSlice := formString(s)
	fmt.Printf("%q\n", stringSlice)
}

func formString(s []string) (ret string) {
	for _, str := range s {
		ret += str
	}
	return
}
