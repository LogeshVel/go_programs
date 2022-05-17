package main

import "fmt"

type money float64

func main() {
	var balance money
	balance = 10.98978
	balance.show()
}

// here we have created the method for money type
// which has float64 as the underlying type
func (m money) show() {
	fmt.Printf("You have $%.2f\n", m)
}
