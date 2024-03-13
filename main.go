package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting port on %s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
