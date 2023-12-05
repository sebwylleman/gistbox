package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from gistbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("starting server on: 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}