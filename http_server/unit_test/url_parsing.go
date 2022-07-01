package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

var persons []Person

func GetFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(persons)
	}
}

func DeleteFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	fmt.Printf("Type:%T,value:%v", person, person)
	for idx, item := range persons {
		if item.ID == person.ID {
			persons = append(persons[:idx], persons[idx+1:]...)
		}
	}
	json.NewEncoder(w).Encode(persons)
}

func UpdateFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	for idx, item := range persons {
		if item.ID == person.ID {
			persons = append(persons[:idx], persons[idx+1:]...)
		}
	}
	persons = append(persons, person)
	json.NewEncoder(w).Encode(person)
}

func CreateFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	persons = append(persons, person)
	json.NewEncoder(w).Encode(persons)
}

func main() {
	persons = append(persons, Person{ID: "2", Name: "Muthu", Age: "22"})
	http.HandleFunc("/get", GetFunc)
	http.HandleFunc("/delete", DeleteFunc)
	http.HandleFunc("/update", UpdateFunc)
	http.HandleFunc("/create", CreateFunc)
	fmt.Println("server stating at port:8081")
	http.ListenAndServe(":8081", nil)
}
