package main

import (
	"fmt"
	"net/http"
)

const PortNum = ":8081"

type person struct {
	name    string
	address string
}

var PersonData person
var StoreData = []person{}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		fmt.Println("GET request called")
		http.ServeFile(w, r, "./home.html")

	case "POST":
		fmt.Println("Post request called")
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		PersonData.name = r.FormValue("name")
		PersonData.address = r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", PersonData.name)
		fmt.Fprintf(w, "Address = %s\n", PersonData.address)
		StoreData = append(StoreData, PersonData)
		fmt.Fprintf(w, "total data:%v", StoreData)

	default:
		fmt.Fprintf(w, "not supported")
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(PortNum, nil)

}
