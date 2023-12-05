package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from gistbox"))
}

func gistView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific gist..."))
}

func gistCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		// Use the http.Error() function to send a 405 status code and
		// "Method not allowed string" as the response body
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new gist..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/gist/view", gistView)
	mux.HandleFunc("/gist/create", gistCreate)

	log.Print("starting server on: 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
