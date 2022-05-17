package main

import "fmt"

type Person struct {
	name   string
	age    int
	status string
}

// func (receiver_type) func_name(){}
func (p Person) print() {
	fmt.Printf("Name :%s\n", p.name)
	fmt.Printf("Age :%d\n", p.age)
	fmt.Printf("Status :%s\n", p.status)
}

type custom_string string

func (some custom_string) print_addr() {
	fmt.Printf("Address of the string is : %p", &some)
}

func main() {
	me := Person{name: "Logesh", age: 23, status: "Unknown"}
	me.print()
	var my_str custom_string
	my_str = "The Logesh Vel"
	my_str.print_addr()
}
