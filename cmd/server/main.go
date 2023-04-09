package main

import (
	"fmt"
	"net/http"
)

func main() {
	// start the web server
	err := http.ListenAndServe(":6161", http.FileServer(http.Dir("assets")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
