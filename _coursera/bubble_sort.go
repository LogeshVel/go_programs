package main

import "fmt"

func BubbleSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				Swap(s, j)
			}
		}
	}
}

func Swap(s []int, i int) {
	s[i], s[i+1] = s[i+1], s[i]
}

func main() {
	var s = []int{}
	for i := 0; i < 10; i++ {
		fmt.Print("Enter slice element ", i+1, ": ")
		var I int
		fmt.Scan(&I)
		s = append(s, I)
	}
	BubbleSort(s)
	fmt.Printf("Sorted list : ")
	fmt.Println(s)
}
