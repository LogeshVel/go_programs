package main

import (
	"fmt"
)

func sayHi(i interface{}) {

	fmt.Println(i, "says hello")

	// Not understanding this error message
	fmt.Println(i.(Dog).name) //  i.name undefined (type interface {} is interface with no methods)
	fmt.Printf("i: %T\n", i)
}

type Dog struct {
	name string
}
type Cat struct {
	name string
}

func main() {
	d := Dog{name: "Sparky"}
	c := Cat{name: "Garfield"}

	sayHi(d) // {Sparky} says hello
	sayHi(c) // {Garfield} says hello
}
