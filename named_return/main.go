package main

import "fmt"

func main() {
	addAndPrint(1, 2)
}

func addAndPrint(a int, b int) (sum int) {
	sum = a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)
	return
}
