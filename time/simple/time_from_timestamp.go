package main

import (
	"fmt"
	"time"
)

func main() {
	var unixTime int64 = 1650872824
	oneSecInNanoSec := 1000000000
	correctTime := time.Unix(unixTime, 0) // 0 nano seconds is added
	fmt.Printf("Correct time : %v\n", correctTime)
	oneSecAddedTime := time.Unix(unixTime, int64(oneSecInNanoSec)) // one second is added
	fmt.Printf("1 sec added  : %v\n", oneSecAddedTime)
}
