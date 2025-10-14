package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles("../../ui/templates/index.html")
	if err != nil {
		fmt.Println("index not found")
		return
	}
	err = templ.Execute(w, nil)
	if err != nil {
		fmt.Println("index not serve")
	}
}
