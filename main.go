package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello and welcome to the GO home page!")
}

func contact(w http.ResponseWriter, r *http.Request) {
	// set the header
	w.Header().Set("Content-Type", "text/html")

	// show the data on the web page
	fmt.Fprintf(w, "Contact us at <a href=\"#\">Hello</a>")
}

func main() {

	// set the variable name for the mux router
	r := mux.NewRouter()

	// create the handler for the home page
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
