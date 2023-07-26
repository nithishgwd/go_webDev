package main

import (
	"net/http"

	"www.goExample.com/nithish/web"
)

func main() {
	http.HandleFunc("/", web.HomePage)
	http.ListenAndServe(":8080", nil)
}
