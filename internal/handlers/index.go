package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"opti-collab/models"
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

func RunCode_handler(w http.ResponseWriter, r *http.Request) {
	var input_code models.Code

	err := json.NewDecoder(r.Body).Decode(&input_code)
	if err != nil {
		fmt.Println("error while decode the given code - ", err)
		return
	}
	fmt.Printf("given code - \n %s - ",input_code.Code)
}
