package main

import (
	"fmt"
	"os"
)

func main() {

	// panic("a problem")
	_, err := os.Create("E:\\_GOlang\\go_programs\\panic\\file.txt")

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfuly created the file")
}
