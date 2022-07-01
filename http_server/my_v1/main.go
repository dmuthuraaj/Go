package main

import (
	"log"
	"net/http"
)

type InMemoryDataStorage struct {
}

func (n *InMemoryDataStorage) GetNumberOfChild(name string) int {
	switch name {
	case "Rajesh":
		return 2
	case "Sajith":
		return 3
	default:
		return http.StatusNotFound
	}
}

func main() {
	server := &DataServer{&InMemoryDataStorage{}}
	log.Fatal(http.ListenAndServe(":8080", server))
}
