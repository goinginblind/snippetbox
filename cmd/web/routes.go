package main

import (
	"net/http"

	"github.com/justinas/alice"
)

// The routes() method returns a handler wrapper for all the other handlers
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// standard var will be a chain of middleware, that wraps around the 'mux'
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	// end up the wrap around and return the whole chain
	return standard.Then(mux)
}
