package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	// Run $ go run ./cmd/web -addr=":PORT NUMBER" with valid port number to set server address.
	// => an idiomatic way of managing configuration settings for this app at runtime
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/gist/view", gistView)
	mux.HandleFunc("/gist/create", gistCreate)

	log.Printf("starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
