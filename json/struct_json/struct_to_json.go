package main

import (
	"encoding/json"
	"fmt"
)

type Candidate struct {
	id   int
	Name string `json:"full_name"`
	Age  int    `json:"age,omitempty"`
}

func main() {

	var c []Candidate
	c = append(c, Candidate{id: 1, Name: "Logesh", Age: 22})
	c = append(c, Candidate{})
	cAsByte, _ := json.Marshal(c)
	fmt.Println(string(cAsByte))
	fmt.Println()
	cAsByte, _ = json.MarshalIndent(c, "", "  ")
	fmt.Println(string(cAsByte))
	fmt.Println()
	cAsByte, _ = json.MarshalIndent(c, "+", "--")
	fmt.Println(string(cAsByte))

}
