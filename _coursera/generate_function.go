package main

import "fmt"

func GenDisplaceFn(a float64, v float64, s float64) func(t float64) float64 {
	return func(t float64) float64 {
		dis := (0.5 * a * t * t) + (v * t) + (s)
		return dis
	}
}

func main() {
	var acc float64
	fmt.Printf("Enter value for acceleration: ")
	fmt.Scan(&acc)

	var iVel float64
	fmt.Printf("Enter value for initial velocity: ")
	fmt.Scan(&iVel)

	var iDis float64
	fmt.Printf("Enter value for initial displacement: ")
	fmt.Scan(&iDis)

	fn := GenDisplaceFn(acc, iVel, iDis)

	var t float64
	fmt.Printf("Enter value for time: ")
	fmt.Scan(&t)

	fmt.Println("Displacement travelled after time", t, "is", fn(t))

}
