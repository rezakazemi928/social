package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct{
	addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", app.healthCheckHandler)
	return r
}

func (app *application) run(mux http.Handler) error{
	srv := http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}
	fmt.Printf("running server on port%v", app.config.addr)
	return srv.ListenAndServe()
}