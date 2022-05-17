package main

import "fmt"

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	var l = Person{Id: 1, Name: "Logesh", Age: 22}
	fmt.Println(l)
}
