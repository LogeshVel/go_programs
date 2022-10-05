package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"encoding/json"
)

func main() {
	var person = map[string]string{}
	fmt.Printf("Enter your name : ")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n') // to read the user input as a sentence. (ex: Logesh Vel)
	line = strings.TrimSpace(line)
	person["name"] = line
	fmt.Printf("Enter your address : ")
	in = bufio.NewReader(os.Stdin)
	line, _ = in.ReadString('\n')
	line = strings.TrimSpace(line)
	person["address"] = line
	pAsByte, _ := json.Marshal(person)
	fmt.Println(string(pAsByte))

}
