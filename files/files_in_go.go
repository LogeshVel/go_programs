package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.ReadFile("E:\\_GOlang\\go_programs\\files\\dummy.txt")
	check(err)
	fmt.Println(string(dat))
	for ln, l := range dat {
		fmt.Println("Line No:", ln, "ASSCII :", l, "Character :", string(l))
	}

	f, err := os.Open("E:\\_GOlang\\go_programs\\files\\dummy.txt")
	check(err)

	b1 := make([]byte, 3)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// o2, err := f.Seek(6, 0)
	// check(err)
	// b2 := make([]byte, 2)
	// n2, err := f.Read(b2)
	// check(err)
	// fmt.Printf("%d bytes @ %d: ", n2, o2)
	// fmt.Printf("%v\n", string(b2[:n2]))

	// o3, err := f.Seek(6, 0)
	// check(err)
	// b3 := make([]byte, 2)
	// n3, err := io.ReadAtLeast(f, b3, 2)
	// check(err)
	// fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// _, err = f.Seek(0, 0)
	// check(err)

	// r4 := bufio.NewReader(f)
	// b4, err := r4.Peek(5)
	// check(err)
	// fmt.Printf("5 bytes: %s\n", string(b4))

	// f.Close()
}