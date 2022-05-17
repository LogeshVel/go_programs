package main

import (
	"fmt"
	"sort"
)

type intSlice []int

func (a intSlice) Len() int           { return len(a) }
func (a intSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intSlice) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	fmt.Println("intSlice")
	a := intSlice{7, 8, 4, 0, 1, 10, 45, 76, 23, 5}
	fmt.Println(a)
	sort.Sort(a)
	fmt.Println(a)

	fmt.Println("intSlice in descending order")
	aa := intSlice{7, 8, 4, 0, 1, 10, 45, 76, 23, 5}
	fmt.Println(aa)
	sort.Sort(sort.Reverse(aa))
	fmt.Println(aa)

	fmt.Println("intSlice([]int)")
	b := []int{7, 8, 4, 0, 1, 10, 45, 76, 23, 5}
	fmt.Println(b)
	sort.Sort(intSlice(b))
	fmt.Println(b)

}
