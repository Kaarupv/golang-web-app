package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
		return
	}
}
