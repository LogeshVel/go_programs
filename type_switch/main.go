package main

import "fmt"

func main() {
	var v interface{}
	fmt.Printf("Type of v (empty interface): %T\n", v)
	v = 10.0
	switch v := v.(type) {
	case int:
		// we can perform any int type operation with v as it is of type int
		fmt.Println("Its int type buddy")
		printValueAndType(v)

	case string:
		fmt.Println("Its string type buddy")
		printValueAndType(v)

	case float64:
		fmt.Println("Its float64 type buddy")
		printValueAndType(v)
	}

}

func printValueAndType(e interface{}) {
	fmt.Printf("Value : %v\n", e)
	fmt.Printf("Type  : %T\n", e)
}
