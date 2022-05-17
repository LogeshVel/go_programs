package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool // #3

type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"` // this will make this field to be skipped in encoding
	Permissions permissions `json:"perms,omitempty"`
	// the omit empty option will make this field not to encode when this field is empty
}

func main() {
	users := []user{ // #2
		{"inanc", "1234", nil}, // here permission field is nill so it wont be encode due to the omitempty option
		{"god", "42", permissions{"admin": true}},
		{"devil", "66", permissions{"write": true}},
	}

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
