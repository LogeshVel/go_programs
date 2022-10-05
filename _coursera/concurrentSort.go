package main

import (
	"fmt"
	"math"
)

func swap(sli []int, i int) {
	temp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}

func bubbleSort(sli []int, c chan []int) { // Bubble sort from course 2, week 1
	changed := true
	for changed {
		changed = false
		for i := 0; i < len(sli)-1; i++ {
			if sli[i] > sli[i+1] {
				swap(sli, i)
				changed = true
			}
		}
	}
	c <- sli
}

func mergeSort(sortedSlis [][]int) []int {
	sorted := make([]int, 0, 20)
	indexes := [4]int{0, 0, 0, 0}
	finished := false
	for !finished {
		min := 10000000000000
		lst := -1
		finished = true
		for i, subSli := range sortedSlis { // Check which list has the current lowest number to be merged.
			if indexes[i] < len(subSli) { // If there are elements in sublist that has not been merged
				if subSli[indexes[i]] < min { // Check if element should be merged
					lst = i
					min = subSli[indexes[i]]
					finished = false
				}
			}
		}
		if !finished {
			sorted = append(sorted, sortedSlis[lst][indexes[lst]]) //Merge the lowest number into sorted list
			indexes[lst] += 1
		}
	}
	return sorted
}

func main() {
	c := make(chan []int)
	sortSli := []int{5, 1, 5, 1, 7654, 432, 123, 654, 86, 51, 5}
	if len(sortSli) < 4 {
		fmt.Println("Slice to be sorted must be longer than four")
		return
	}
	fmt.Println("Unsorted list", sortSli)

	subLen := math.Ceil(float64(len(sortSli)) / 4) // Length of sub lists
	for i := 0; i < 3; i++ {
		subSli := sortSli[i*int(subLen) : (i+1)*int(subLen)]
		go bubbleSort(subSli, c) // Sort sub lists
	}
	subSli := sortSli[3*int(subLen) : len(sortSli)] // If nr of lists not a multiple of 4
	go bubbleSort(subSli, c)                        // Sort sub lists
	sortedSlis := make([][]int, 0, 0)
	for i := 0; i < 4; i++ {
		sortedSli := <-c // Gather all 4 sub lists
		sortedSlis = append(sortedSlis, sortedSli)
	}
	fmt.Println("Sorted sub lists", sortedSlis)

	fmt.Println("Sorted list", mergeSort(sortedSlis)) // Merge sort on sorted sub lists
}
