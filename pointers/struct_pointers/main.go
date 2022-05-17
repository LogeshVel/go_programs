package main

import "fmt"

type house struct {
	name  string
	rooms int
}

func main() {
	structs()
}

func structs() {
	myHouse := house{name: "My House", rooms: 5}

	addRoom(myHouse)

	// fmt.Printf("%+v\n", myHouse)
	fmt.Printf("structs()     : %p %+v\n", &myHouse, myHouse)

	addRoomPtr(&myHouse)
	fmt.Printf("structs()     : %p %+v\n", &myHouse, myHouse)
}

func addRoomPtr(h *house) {
	h.rooms++ // same: (*h).rooms++
	fmt.Printf("addRoomPtr()  : %p %+v\n", h, h)
	fmt.Printf("&h.name       : %p\n", &h.name)
	fmt.Printf("&h.rooms      : %p\n", &h.rooms)
}

func addRoom(h house) {
	h.rooms++
	fmt.Printf("addRoom()     : %p %+v\n", &h, h)
}
