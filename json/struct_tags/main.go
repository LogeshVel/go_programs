package main

import (
	"encoding/json"
	"fmt"
)

type NormalPerson struct {
	Id   int
	Name string
	Age  int
}

type TagPerson struct {
	Id   int    `json:"id"`
	Name string `json:"full_name"`
	Age  int    `json:"age"`
}

func main() {
	var a NormalPerson
	aAsByte, _ := json.MarshalIndent(a, "", "  ")
	fmt.Println("Normal Person")
	fmt.Println(string(aAsByte))

	var b TagPerson
	bAsByte, _ := json.MarshalIndent(b, "", "  ")
	fmt.Println("Tag Person")
	fmt.Println(string(bAsByte))
}
