package main

import "fmt"

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	var l = Person{Id: 1, Name: "Logesh", Age: 22}
	// fmt.Println(l)
	fmt.Printf("Type of &l : %T\n", &l)
	fmt.Printf("Type of  l : %T\n", l)

}

func (p *Person) String() string {
	return fmt.Sprintf("Person Type\nId   : %d\nName : %s\nAge  : %d", p.Id, p.Name, p.Age)
}
