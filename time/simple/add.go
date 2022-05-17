package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now()
	fmt.Printf("\nCurrent time   :\t%v\n", nowTime)
	fmt.Printf("\nCurrent Second :\t%v\n\n", nowTime.Second())

	nowTime = nowTime.Add(2 * time.Second)
	fmt.Println("After adding 2 seconds")
	fmt.Printf("\nCurrent Second :\t%v\n", nowTime.Second())

}
