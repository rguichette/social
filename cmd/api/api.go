package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config serverConfig
}

type serverConfig struct {
	addr string
}

func (app *application) mount() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

	return mux

}
func (a *application) start(mux *http.ServeMux) error {
	srv := http.Server{
		Addr:         a.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s", a.config.addr)

	return srv.ListenAndServe()

}
