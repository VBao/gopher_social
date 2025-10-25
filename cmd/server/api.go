package main

import (
	"log"
	"net/http"
	"time"
)

// Use to inject all the required service, config for the application
type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", app.healthCheckHandler)

	return mux
}

func (app *application) run(mux *http.ServeMux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30, // Server handle request time out
		ReadTimeout:  time.Second * 10, // Client read request time out
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server run at %s", app.config.addr)

	return srv.ListenAndServe()
}
