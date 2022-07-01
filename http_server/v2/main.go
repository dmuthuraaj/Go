package main

import (
	"log"
	"net/http"
)

//collecting a name data
type InMemoryNameStore struct {
}

func (i *InMemoryNameStore) GetPersonAge(name string) int {
	return 23
}

func main() {
	server := &GetNameServer{&InMemoryNameStore{}}  //this creates a handler interface
	log.Fatal(http.ListenAndServe(":8080", server)) //starting server at port:8080
}
