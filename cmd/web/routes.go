package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileSever := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileSever))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /jobapp/view/{id}", app.jobAppView)
	mux.HandleFunc("GET /jobapp/create", app.jobAppCreate)
	mux.HandleFunc("POST /jobapp/create", app.jobAppCreatePost)

	return mux
}
