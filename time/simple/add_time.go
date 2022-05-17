package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now()
	fmt.Printf("\nCurrent time :\t%v\n", nowTime)
	nowTime = nowTime.Add(time.Hour)
	fmt.Println("After adding one hour")
	fmt.Printf("\t\t%v\n", nowTime)

	fmt.Printf("\nSome time constants\n\n")
	fmt.Printf("time.Hour        %20d\n", time.Hour)
	fmt.Printf("time.Minute      %20d\n", time.Minute)
	fmt.Printf("time.Second      %20d\n", time.Second)
	fmt.Printf("time.Millisecond %20d\n", time.Millisecond)
	fmt.Printf("time.Microsecond %20d\n", time.Microsecond)
	fmt.Printf("time.Nanosecond  %20d\n", time.Nanosecond)

}
