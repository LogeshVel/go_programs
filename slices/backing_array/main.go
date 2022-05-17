package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4, 5}

	s2 := s1[0:2]

	fmt.Printf("len(s2) : %d\t cap(s2) : %d\n", len(s2), cap(s2))
	fmt.Printf("s2: %#v\n", s2)
	fmt.Printf("Getting 4th index from s2: %d\n", s2[4:5])
	fmt.Printf("len(s2) : %d\t cap(s2) : %d\n", len(s2), cap(s2))
}
