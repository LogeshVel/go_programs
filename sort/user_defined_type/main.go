package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type personByAge []Person

func (a personByAge) Len() int           { return len(a) }
func (a personByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a personByAge) Less(i, j int) bool { return a[i].age < a[j].age }

type personByName []Person

func (a personByName) Len() int           { return len(a) }
func (a personByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a personByName) Less(i, j int) bool { return a[i].name < a[j].name }

func main() {
	p1 := Person{name: "Logesh", age: 100}
	p2 := Person{name: "Vel", age: 101}
	p3 := Person{name: "Sam", age: 18}
	p4 := Person{name: "Vikram", age: 78}
	p5 := Person{name: "Aakash", age: 42}
	p6 := Person{name: "Ram", age: 31}

	p1Slice := []Person{p1, p2, p3, p4, p5, p6}
	p2Slice := []Person{p1, p2, p3, p4, p5, p6}

	fmt.Println("Sorting Person Slice by Age")
	fmt.Printf("Before : %v\n", p1Slice)
	sort.Sort(personByAge(p1Slice))
	fmt.Printf("After  : %v\n", p1Slice)

	fmt.Println("Sorting Person Slice by Name")
	fmt.Printf("Before : %v\n", p2Slice)
	sort.Sort(personByName(p2Slice))
	fmt.Printf("After  : %v\n", p2Slice)

}
