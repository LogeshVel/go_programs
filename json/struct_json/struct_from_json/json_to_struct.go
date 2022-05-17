package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Candidate struct {
	id   int
	Name string `json:"full_name"`
	Age  int    `json:"age"`
}

const jsonData = `
[
  {
    "id": 1,
    "full_name": "Logesh",
    "age": 22
  },
  {
    "id": 0,
    "full_name": "",
    "age": 0
  }
]
`

func main() {

	var c []Candidate
	printCandidateList(c)
	err := json.Unmarshal([]byte(jsonData), &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nAfter storing the data\n\n")
	printCandidateList(c)

}

func printCandidateList(cList []Candidate) {
	if len(cList) == 0 {
		fmt.Println("No Candidate is added to the list")
		return
	}
	for i, c := range cList {
		fmt.Printf("Candidate #%v : %+v\n", i+1, c)
	}
}
