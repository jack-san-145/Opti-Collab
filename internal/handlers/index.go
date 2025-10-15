package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"opti-collab/docker"
	"opti-collab/internal/services"

	// "opti-collab/internal/services"
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
	fmt.Println("given language - ", input_code.Language)
	fmt.Printf("given code - \n %s - ", input_code.Code)

	code_output, err := docker.Run_code(input_code.Language, input_code.Code)
	if err != nil {
		fmt.Println("error while runnnig the code in docker - ", err)
	} else {
		WriteJSON(w, r, map[string]string{"output": code_output})
	}

}

func FindOptmiseCode_handler(w http.ResponseWriter, r *http.Request) {
	var OptimiseCode models.Code
	err := json.NewDecoder(r.Body).Decode(&OptimiseCode)
	if err != nil {
		fmt.Println("error while decode the given code - ", err)
		return
	}
	response, err := services.AnalyzeCode(OptimiseCode.Code, OptimiseCode.Language)
	if err != nil {
		fmt.Println("error while find the optimised code - ", err.Error())
	}
	WriteJSON(w, r, response)
}
