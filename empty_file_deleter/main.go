package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	flagGivenPath := flag.String("p", "", "Path, from where the empty files need to find and proceed for deletion")

	givenPath := *flagGivenPath
	if givenPath == "" {
		fmt.Println("Provide the Path to find the empty files")
		return
	}
	files, err := ioutil.ReadDir(givenPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	var emptyFileCount int
	// this optimization wastes cpu time but memory efficient
	for _, f := range files {
		if f.Size() == 0 {
			emptyFileCount += len(f.Name()) + 1 // here +1 is due to the new char i going to add
		}
	}
	firstLine := fmt.Sprintf("Empty files name list under the path : %s\n", givenPath)
	emptyFileCount += len(firstLine)
	emptyFileNames := make([]byte, 0, emptyFileCount)
	fmt.Printf("Size of emptyFileNames: %v\n", emptyFileCount)
	emptyFileNames = append(emptyFileNames, firstLine...)
	for _, file := range files {
		f_name := file.Name()
		if file.Size() == 0 {
			fmt.Printf("Found Empty file: %s\n", f_name)
			emptyFileNames = append(emptyFileNames, f_name...)
			emptyFileNames = append(emptyFileNames, '\n')
		}
	}
	errr := ioutil.WriteFile("empty_filesname.txt", emptyFileNames, 0644)
	if errr != nil {
		fmt.Println(errr)
		return
	}
}
