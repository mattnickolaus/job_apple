package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()

	fileSever := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileSever))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /jobapp/view/{id}", app.jobAppView)
	mux.HandleFunc("GET /jobapp/create", app.jobAppCreate)
	mux.HandleFunc("POST /jobapp/create", app.jobAppCreatePost)

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)

	logger.Error(err.Error())
	os.Exit(1)
}
