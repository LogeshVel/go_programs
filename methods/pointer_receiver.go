package main

import (
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	l := Person{Name: "Logesh", Age: 22}
	l.print()

	l.up()
	// (&l).up()

	l.print()

}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}

// Pointer receiver to modify the values
func (p *Person) up() {
	p.Name = strings.ToUpper(p.Name)
}
