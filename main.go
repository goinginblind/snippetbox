package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Home page handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from the SnippetBox!\n"))
}

// Provides a page with a specific snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Displaying a specific snippet with an id %d...\n", id)
	w.Write([]byte(msg))
}

// Provides a page with a snippet creation page: a GET request must be sent
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here a snippet would've been created\n"))
}

// Creates a snippet: a POST request must be sent
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Here a snippet would've been saved\n"))
}

// For now, the main entry point of the whole app
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("Starting the server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
