package main

import (
	"fmt"
	"net/http"
)

// Handles the root page requests.
func rootPage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("\nHitted rootPage")
}

// Handles the echo page requests.
func echoPage(w http.ResponseWriter, request *http.Request) {
	fmt.Println("\nHitted echoPage")
}

func main() {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/echo", echoPage)

	http.ListenAndServe(":8090", nil)

}
