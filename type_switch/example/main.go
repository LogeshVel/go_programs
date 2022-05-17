package main

import (
	"fmt"
)

func main() {
	var v interface{}
	fmt.Printf("Type of v (empty interface): ")

	_ = v

	// v = v.(type) // error

}
