package main

import (
	"encoding/json"
	"fmt"
)

type TagPerson struct {
	Id   int    `json:"id"`
	Name string `json:"full_name"`
	Age  int    `json:",omitempty"`
}

func main() {

	var b TagPerson
	printInJSON(b)
	b.Age = 21
	printInJSON(b)

}

func printInJSON(tp TagPerson) {
	bAsByte, _ := json.MarshalIndent(tp, "", "  ")
	fmt.Println(string(bAsByte))
}
