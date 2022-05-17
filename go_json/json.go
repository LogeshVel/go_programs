package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	mapD := map[string]int{"one": 5, "two": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.SetIndent("", "  ")
	enc.Encode(d)
	enc.Encode(mapD)
}
