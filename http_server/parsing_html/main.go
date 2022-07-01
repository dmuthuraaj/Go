package main

import (
	"net/http"
)

const PORTNUMBER = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(PORTNUMBER, nil)
}
