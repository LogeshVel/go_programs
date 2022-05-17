package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{name: "Logesh", age: 100}
	p2 := Person{name: "Vel", age: 101}
	p3 := Person{name: "Sam", age: 18}
	p4 := Person{name: "Vikram", age: 78}
	p5 := Person{name: "Aakash", age: 42}
	p6 := Person{name: "Ram", age: 31}

	p1Slice := []Person{p1, p2, p3, p4, p5, p6}
	// fmt.Println("Sorting by age")
	// fmt.Println(p1Slice)
	// sort.Slice(p1Slice, func(i, j int) bool {
	// 	return p1Slice[i].age < p1Slice[j].age
	// })
	fmt.Println(p1Slice)
	fmt.Println("Sorting by name")
	fmt.Println(p1Slice)
	sort.SliceStable(p1Slice, func(i, j int) bool {
		return p1Slice[i].name < p1Slice[j].name
	})
	fmt.Println(p1Slice)

}
