package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const loop = 100
	var wg sync.WaitGroup
	wg.Add(loop * 2)

	// declaring a shared value
	var n int = 0

	for i := 0; i < loop; i++ {
		// one go routine increments the value
		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()
		// second go routine decrement the value
		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
	}
	wg.Wait()
	// printing the final value of n
	if n != 0 {
		fmt.Println("The Final value of n should be 0. But found ", n)
		return
	}
	fmt.Printf("\nFinal value of n is %d\n\n", n) // the final of n should be 0
}

// Here in the given program I have made 2 go routines to access the shared variable.

// One Go routine will try to increment the value of the variable
// Another Go routine will try to decrement the value of the variable.

// so at the end the value of the varibale should be same as the initialization value. But due to the race condition the value won't be the same as initialization value all time.
