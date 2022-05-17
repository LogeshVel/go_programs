package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println
	p(time.Now())
	pt, _ := time.Parse("2006-01-2 15:04:05.000000", "2021-12-19 20:23:27.962903")
	p(pt.Unix())
}
