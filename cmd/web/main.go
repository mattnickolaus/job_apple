package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	fileSever := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileSever))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /jobapp/view/{id}", jobAppView)
	mux.HandleFunc("GET /jobapp/create", jobAppCreate)
	mux.HandleFunc("POST /jobapp/create", jobAppCreatePost)

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)

	logger.Error(err.Error())
	os.Exit(1)
}
