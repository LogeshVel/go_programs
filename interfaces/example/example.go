package main

import (
	"fmt"
)

// declaring an interface type called calculation
// an interface contains only the signatures of the methods, but not their implementation
type calculation interface {
	add() float64
	mul() float64
}

// declaring 2 struct types

type two_inputs struct {
	a, b float64
}
type variable_inputs struct {
	list []float64
}

// method that calculates addition of two numbers
func (t two_inputs) add() float64 {
	return t.a + t.b
}

// method that calculates addition of slice
func (v variable_inputs) add() float64 {
	s := 0.0
	for _, e := range v.list {
		s += e
	}
	return s
}

// method that calculates multiplication of two numbers
func (t two_inputs) mul() float64 {
	return t.a * t.b
}

// method that calculates multiplication of elements in the slice
func (v variable_inputs) mul() float64 {
	m := 1.0
	for _, e := range v.list {
		m *= e
	}
	return m
}

// any type that implements the interface is also of type of the interface
// two_inputs and variable_inputs values are also of type calculation
func get_calcs(c calculation) {
	fmt.Printf("Calculation type : %#v\n", c)
	fmt.Printf("Add : %v\n", c.add())
	fmt.Printf("Mul : %v\n", c.mul())
}

func main() {

	// declaring a two_inputs and a variable_inputs types
	t1 := two_inputs{a: 2.0, b: 8.7}
	s1 := variable_inputs{list: []float64{1.0, 2.5, 9.1}}

	// two_inputs and variable_inputs both implements the calculation interface
	// because they implement all methods of the interface
	// an interface is implicitly implemented in Go
	get_calcs(t1)
	get_calcs(s1)
}
