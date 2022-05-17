package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	new_file, err := os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(new_file, "Logesh Vel")
	created_file, err := os.ReadFile("a.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(created_file))
	fmt.Fprintln(new_file, "Again Logesh Vel")
	created_file, err = os.ReadFile("a.txt")
	_ = err
	fmt.Println(string(created_file))
}

func file_works() {

	//** Use valid paths according to your OS. **//

	// CREATING A FILE

	// os.Create() function creates a file if it doesn't already exist. If it exists, the file is truncated.
	// it returns a file descriptor which is a pointer to os.File and an error value.
	newFile, err := os.Create("a.txt")

	// error handling
	if err != nil {
		// log the error and exit the program
		log.Fatal(err) // the idiomatic way to handle errors

	}

	fmt.Fprintln(newFile, "written in program by Logesh!!!")

	// TRUNCATING A FILE
	err = os.Truncate("a.txt", 0) //0 means completely empty the file.

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(newFile, "written in program by Logesh!!!, its erased and now written")

	// CLOSING THE FILE
	newFile.Close()

	// OPEN AND CLOSE AN EXISTING FILE
	file, err := os.Open("a.txt") // open in read-only mode

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	var f_r = []byte{10}
	contents, err := file.Read(f_r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n")
	fmt.Println("Contents of the file", file.Name())
	fmt.Println(contents)
	fmt.Printf("\n")

	file.Close()

	//OPENING a FILE WITH MORE OPTIONS
	file, err = os.OpenFile("a.txt", os.O_APPEND, 0644)
	// We can Use opening attributes individually or combined
	// using an OR between them
	// e.g. os.O_CREATE|os.O_APPEND
	// or os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	// os.O_RDONLY // Read only
	// os.O_WRONLY // Write only
	// os.O_RDWR // Read and write
	// os.O_APPEND // Append to end of file
	// os.O_CREATE // Create is none exist
	// os.O_TRUNC // Truncate file when opening

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// GETTING FILE INFO
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	p := fmt.Println
	p("File Name:", fileInfo.Name())        // => File Name: a.txt
	p("Size in bytes:", fileInfo.Size())    // => Size in bytes: 0
	p("Last modified:", fileInfo.ModTime()) // => Last modified: 2019-10-21 16:16:00.325037748 +0300 EEST
	p("Is Directory? ", fileInfo.IsDir())   // => Is Directory?  false
	p("Pemissions:", fileInfo.Mode())       // => Pemissions: -rw-r-----

	// CHECKING IF FILE EXISTS
	fileInfo, err = os.Stat("b.txt")
	// error handling
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist")
		}
	}
	_ = fileInfo
	// RENAMING AND MOVING A FILE
	oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	// error handling
	if err != nil {
		log.Fatal(err)
	}

	// REMOVING A FILE
	err = os.Remove("aa.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}
