package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, "about.html")
}

func RenderPage(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./template/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error passing template", err)
		return
	}

}
