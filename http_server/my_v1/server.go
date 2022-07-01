package main

import (
	"fmt"
	"net/http"
	"strings"
)

type DataStore interface {
	GetNumberOfChild(name string) int
}

type DataServer struct {
	store DataStore
}

func (s *DataServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/name/")
	NumberOfChild := s.store.GetNumberOfChild(name)
	if NumberOfChild == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, NumberOfChild)
}
