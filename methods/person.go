package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	l := Person{Name: "Logesh", Age: 22}
	l.print()
	printPerson(l)

}

// this is method syntax
// func (receiver receiverType) methodname(){}
// this method belongs to that receiver type
// only that particular type can use this method
func (p Person) print() {
	fmt.Printf("%#v\n", p)
}

// this is function
func printPerson(p Person) {
	fmt.Printf("%+#v\n", p)
}
