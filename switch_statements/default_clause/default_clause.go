package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ar := os.Args[1]
	a, _ := strconv.Atoi(ar)
	// at the end of each case there is a virtual break
	// so its not trying to execute next case like in other langs
	switch a {
	// the default clause can be in anywhere in the switch statement.
	// it will be executed at the end if no case is statisfies the condition
	default:
		fmt.Println("a is", a)
	case 2, 4: // multiple conditions(this case will be executed if a is either 2 or 4)
		fmt.Println("a is 2, i found it")
	}
}
