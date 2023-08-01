package main

import (
	"fmt"
	"net/http"

	"www.goExample.com/nithish/web"
)

const (
	portNumber = 8080
)

func main() {

	// Started the server with the chosen port number
	addr := fmt.Sprintf(":%d", portNumber)

	fmt.Printf("Application is running on localhost%v\n", addr)

	// Set up the HTTP routes and handlers
	http.HandleFunc("/", web.HomePage)

	http.HandleFunc("/contact", web.ContactPage)

	http.ListenAndServe(addr, nil)

}
