package main

import "fmt"

func changer(a ...int) {
	a[0] = 109
	a[1] = 456
}

func plus_changer(b int, others ...string) {
	b = 100001
	others[0] = "chnged"
}

func main() {

	my_slice := []int{1, 2, 3, 4, 5}
	fmt.Println(my_slice)
	changer(my_slice...)
	fmt.Println(my_slice)

	my_strs := []string{"Logesh", "vel"}
	my_var := 10
	fmt.Println(my_var, my_strs)
	plus_changer(my_var, my_strs...)
	fmt.Println(my_var, my_strs)

}
