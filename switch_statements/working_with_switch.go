package main

import "fmt"

func main() {
	fmt.Println("Working with switch, with expression")
	a := 10
	// at the end of each case there is a virtual break
	// so its not trying to execute next case like in other langs
	switch a {
	case 2:
		fmt.Println("a is 2, i found it")
	case 10:
		fmt.Println("a is 10, i found it man")
	default:
		fmt.Println("a is", a)
	}
	fmt.Println("Now working with switch without expression")
	// if the case condtion is true then that case will execute
	switch {
	case a == 2:
		fmt.Println("a is 2, i found it")
	case a == 10:
		fmt.Println("a is 10, i found it man")
	default:
		fmt.Println("a is", a)
	}

	fmt.Println("Now working with switch true expression")
	// if the case condtion is true then that case will execute
	switch true {
	case a == 2:
		fmt.Println("a is 2, i found it")
	case a == 10:
		fmt.Println("a is 10, i found it man")
	default:
		fmt.Println("a is", a)
	}

	fmt.Println("Now working with switch false expression")
	// if the case condtion is false then that case will execute
	switch false {
	case a == 2:
		fmt.Println("a is 2, i found it")
	case a == 10:
		fmt.Println("a is 10, i found it man")
	default:
		fmt.Println("a is", a)
	}
}
