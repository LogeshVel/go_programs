package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now()
	fmt.Println(nowTime)
	fmt.Printf("Unix time  : %v\n", nowTime.Unix())
	fmt.Printf("UTC  time  : %v\n", nowTime.UTC())
	fmt.Printf("Time in str: %v its type '%T'\n", nowTime.GoString(), nowTime.GoString())

}
