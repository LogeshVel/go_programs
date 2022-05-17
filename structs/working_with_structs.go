package main

import "fmt"

type Person struct {
	name   string
	age    int
	status string
}

func main() {
	var p1 Person
	p1.name = "Logesh"
	p1.age = 23
	p1.status = "Single"
	fmt.Println(p1)
	fmt.Printf("%v\n", p1)
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%+v\n", p1)

	// var p2 = Person{}
	// fmt.Printf("p2: %#v\n", p2)

	// p3 := Person{name: "Someone", age: 3422, status: "Not known"}
	// fmt.Printf("p3: %#v\n", p3)
}
