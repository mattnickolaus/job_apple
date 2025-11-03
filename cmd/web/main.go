package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello there! Welcome to Job Apple"))
}

func jobAppView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display a specific job applicaiton with ID %d", id)
	w.Write([]byte(msg))
}

func jobAppCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new job application..."))
}

func jobAppCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("Save a new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /jobapp/view/{id}", jobAppView)
	mux.HandleFunc("GET /jobapp/create", jobAppCreate)
	mux.HandleFunc("POST /jobapp/create", jobAppCreatePost)

	log.Printf("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
