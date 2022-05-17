package main

import (
	"fmt"
)

func main() {
	var s1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var s2 = []int{11, 12, 13, 7, 2, 8}
	var loopFlag bool
	for _, ele1 := range s1 {
		for _, ele2 := range s2 {
			if ele1 == ele2 {
				loopFlag = true
				fmt.Printf("Match elements %d : %d\n", ele1, ele2)
				break
			}
		}
		if loopFlag {
			break
		}
	}

}
