package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	env  string
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "port where application will be run")
	flag.StringVar(&cfg.env, "env", "development", "current environment of the applicaton")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)
	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
