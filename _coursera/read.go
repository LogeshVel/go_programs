package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type names struct{
	fname string
	lname string
}

func main() {
	namesSlice := []names{}
	var filename string
	fmt.Printf("Enter the filename/filepath: ")
	fmt.Scan(&filename)
	readFile, err := os.Open(filename)
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
		line := fileScanner.Text()
		words := strings.Fields(line)
        f := words[0]
        l := words[1]
		namesSlice = append(namesSlice, names{fname: f, lname: l})
    }

	for _, p := range(namesSlice){
		fmt.Println("First name : ",p.fname, "\t", "Last name : ", p.lname)
	}
  
    readFile.Close()


}