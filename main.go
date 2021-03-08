package main

import (
	"fmt"
	"net/http"
)

func homeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello and welcome to the GO home page!")
}

func main() {
	// create the handler for the home page
	http.HandleFunc("/", homeFunc)
	http.ListenAndServe(":3000", nil)
}
