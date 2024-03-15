package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting port on %s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
