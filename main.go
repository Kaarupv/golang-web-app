package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const PORT = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
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

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting port on %s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
