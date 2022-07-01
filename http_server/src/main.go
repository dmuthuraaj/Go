package main

import (
	"handlers"
	"net/http"
)

const PORTNUMBER = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(PORTNUMBER, nil)
}
