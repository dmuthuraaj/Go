package handlers

import (
	"myapp/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, "about.page.html")
}
