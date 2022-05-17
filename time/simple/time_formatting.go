package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now()

	fmt.Println(nowTime)

	// The layout string used by the Parse function and Format method
	// shows by example how the reference time should be represented.
	// We stress that one must show how the reference time is formatted,
	// not a time of the user's choosing. Thus each layout string is a
	// representation of the time stamp,
	//	Jan 2 15:04:05 2006 MST

	// here we can have any type of format but
	// we need to use the reference time in our format
	// here i need my time in the format of
	// year % month % date
	layout := "2006 % Jan % 2"
	t := nowTime.Format(layout)
	fmt.Println("Formatted time :", t)

}
