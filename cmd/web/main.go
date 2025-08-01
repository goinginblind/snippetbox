package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Holds dependencies passed into the handlers, its more of a placeholder as of now.
// Holds only one piece of data: logger
type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network adress")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	logger.Info("starting server", slog.String("addr", *addr))

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())

	os.Exit(1)
}
