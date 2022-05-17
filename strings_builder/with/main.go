package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = []string{"One", "Two", "Three", "Four", "Five"}
	stringSlice := formStringAndReste(s)
	fmt.Printf("\n%q\n", stringSlice)
	formString(s)
}

func formStringAndReste(s []string) (result string) {
	var sb strings.Builder
	fmt.Printf("Initial %25s Len : %3d, Cap : %3d, ptr : %3p\n", " ", sb.Len(), sb.Cap(), &sb)
	for _, str := range s {
		sb.WriteString(str)
		fmt.Printf("String : %5s, String Len : %3d, Len : %3d, Cap : %3d, ptr : %3p\n", str, len(str), sb.Len(), sb.Cap(), &sb)

	}
	result = sb.String()
	sb.Reset()
	fmt.Printf("After Reset %21s Len : %3d, Cap : %3d, ptr : %3p\n", " ", sb.Len(), sb.Cap(), &sb)

	return
}

func formString(s []string) string {
	var sb strings.Builder
	fmt.Printf("Initial %25s Len : %3d, Cap : %3d, ptr : %3p\n", " ", sb.Len(), sb.Cap(), &sb)
	for _, str := range s {
		sb.WriteString(str)
		fmt.Printf("String : %5s, String Len : %3d, Len : %3d, Cap : %3d, ptr : %3p\n", str, len(str), sb.Len(), sb.Cap(), &sb)

	}

	return sb.String()
}
