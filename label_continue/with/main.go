package main

import (
	"fmt"
)

func main() {
	var s1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var s2 = []int{11, 12, 13, 7, 2, 8}
	var printTimes int = 2

	for _, ele1 := range s1 {
	secondloop:
		for _, ele2 := range s2 {
			if ele2 == ele1 {
				for i := 0; i < printTimes; i++ {
					if ele1%2 == 0 {
						fmt.Printf("Matched elements : %d = %d : Even\n", ele1, ele2)
						continue secondloop
					} else {
						fmt.Printf("Matched elements : %d = %d : Odd\n", ele1, ele2)
					}
				}
			}
		}

	}

}
