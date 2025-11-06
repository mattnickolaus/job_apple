package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileSever := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileSever))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /jobapp/view/{id}", jobAppView)
	mux.HandleFunc("GET /jobapp/create", jobAppCreate)
	mux.HandleFunc("POST /jobapp/create", jobAppCreatePost)

	log.Printf("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
