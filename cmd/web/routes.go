package main

import "net/http"

// the routes() method returns a servemux containing application routes
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/gist/view", app.gistView)
	mux.HandleFunc("/gist/create", app.gistCreate)

	return mux
}
