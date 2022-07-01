package main

import (
	"fmt"
	"net/http"
	"strings"
)

type NameStore interface {
	GetPersonAge(name string) int
}

type GetNameServer struct {
	store NameStore
}

func (p *GetNameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/name/")
	age := p.store.GetPersonAge(name)
	if age == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, age)
}
