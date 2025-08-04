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

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	// standard var will be a chain of middleware, that wraps around the 'mux'
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	// end up the wrap around and return the whole chain
	return standard.Then(mux)
}
