package render

import (
	"fmt"
	"html/template"
	"http/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderPage(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./template/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error passing template", err)
		return
	}

}

func RenderPageTest(w http.ResponseWriter, tmpl string) error {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./template/*.page.html")
	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).funcs(functions).Parsefiles(page)
		if err != nil {
			return err
		}

	}
}
