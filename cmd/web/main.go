package main

// TODO - Allow user to sign up to DailyCents
// TODO - Allow user to set up their banking info with Plaid
// TODO - Allow user to set up notifications with AWS SNS

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
	// Grab port number from command line argument or use default value of :2000
	addr := flag.String("addr", ":2000", "HTTP network address")
	flag.Parse()

	// Inizialize new JSON logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Inizialize new application struct
	app := application{
		logger: logger,
	}

	// Inizialize new customer http.Server struct
	srv := http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", slog.String("addr", *addr))

	// Start server, log and exit if any errors
	err := srv.ListenAndServe()

	logger.Error(err.Error())
	os.Exit(1)
}
