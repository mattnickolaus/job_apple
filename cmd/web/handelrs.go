package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type JobApplication struct {
	Company     string
	Title       string
	DateApplied time.Time
	Link        string
	Status      string
	LastUpdated time.Time
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := []JobApplication{
		{"NASA", "Software Engineer 1", time.Now(), "https://nasa.gov", "Submitted", time.Now()},
		{"NASA", "Software Engineer 1", time.Now(), "https://nasa.gov", "Submitted", time.Now()},
		{"NASA", "Software Engineer 1", time.Now(), "https://nasa.gov", "Submitted", time.Now()},
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func jobAppView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func jobAppCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new job application..."))
}

func jobAppCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet"))
}
