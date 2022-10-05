package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	intSlice := make([]int, 3)
	for {
		var s string
		var breaker = "X"
		fmt.Printf("Enter an int : ")
		fmt.Scan(&s)
		if s == breaker {
			fmt.Println("Exiting!")
			return
		}
		i, err := strconv.Atoi(s)
		if err != nil{
			continue
		}
		intSlice = append(intSlice, i)
		sort.Slice(intSlice, func(i, j int) bool {
			return intSlice[i] < intSlice[j]
		})
		fmt.Printf("The Sorted slice is : %v\n", intSlice)

	}

}
