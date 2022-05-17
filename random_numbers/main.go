package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	// everytime i ran this code it will give the same results,
	// which says the rand is not truly random

	// to make it truly random we need to randomize the rand seed
	rand.Seed(time.Now().UnixNano())
	// time is the only thing that is unique all time. so by giving the unique seed value
	// everytime the rand will generate the true random numbers
	fmt.Print("\nAfter seeded\n")
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

}
