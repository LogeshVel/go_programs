package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4, 5}
	fmt.Printf("s1          : %#v\n", s1)
	s1 = append(s1, s1[2:]...)
	fmt.Printf("s1          : %#v\n", s1)
	s1 = append(s1[:2], []int{9, 16, 25}...)
	fmt.Printf("s1          : %#v\n", s1)
	fmt.Printf("s1[:cap(s1)]: %#v\n", s1[:cap(s1)])

}
