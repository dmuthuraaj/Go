package main

import (
	"fmt"
	"net/http"
)

func GetName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dhana")
}
