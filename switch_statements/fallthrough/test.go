package main

import "fmt"

func main() {
	i := 1
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i > 50: // this confirms that the fallthrough statement does not validate the next case clause expression
		fmt.Println("i is less than 50")
		fallthrough
	case i < 100:
		fmt.Println("i is less than 100")
	}
}

// if this line is executed then only the next case block is executed without validating the expression
