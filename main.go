package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParsedForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	// these values come from the form that is submitted in form.html
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// * is a pointer and r points to this Request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// we don't want any post, update, delete requests for this route
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	// := short form operator that declares and defines a variable
	// tell golang to look at the static directory. Golang will automatically search for an index.html by default
	fileServer := http.FileServer(http.Dir("./static"))
	// handle the root route
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	// this line will create the server. The
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
