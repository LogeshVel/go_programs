package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sortSubArray(wg *sync.WaitGroup, i, subArray int, arr []int) {
	start := int(math.Min(float64(i*subArray), float64(len(arr))))
	end := int(math.Min(float64((i+1)*subArray), float64(len(arr))))
	subArr := arr[start:end]
	fmt.Println("Goroutine ", i, " : will sort this Array : ", subArr)
	sort.Ints(subArr)
	fmt.Println("Goroutine ", i, " : result : ", subArr)
	wg.Done()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter a list of integers : ")
	fmt.Println("(example of list : 1 8 7 2 3 6 5 4 )")
	a, _ := reader.ReadString('\n')

	s := strings.Split(strings.TrimSpace(a), " ")

	var arr []int
	for _, str := range s {
		n, _ := strconv.Atoi(str)
		arr = append(arr, n)
	}

	subArray := int(math.Max(math.Ceil(float64(len(s))/float64(4)), 1))

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go sortSubArray(&wg, i, subArray, arr)
	}

	wg.Wait()
	sort.Ints(arr)
	fmt.Println("Sorted Array:", arr)
}
