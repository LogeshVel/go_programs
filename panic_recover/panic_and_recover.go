package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("By the defer")
		err := recover()
		if err != nil {
			fmt.Println("is errored so trying again")
		}
		fmt.Println(err)
	}()
	fmt.Println("Before panic")
	// panic("purpose")
	fmt.Println("after panic")
}
