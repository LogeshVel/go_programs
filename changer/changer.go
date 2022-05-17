package main

import "fmt"

func changer(list *[]int) {
	for i := 0; i < 10; i++ {
		*list = append(*list, i+1)

	}
	fmt.Println("List in function ", *list)
}

func main() {
	// list is modified when we pass by refrence (address)
	var a []int
	fmt.Println("emp:", a)
	changer(&a)
	fmt.Println("emp:", a)

	// we don't need pass by reference for map
	// pass by value behaves the same

}
