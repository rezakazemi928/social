package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Health check passed\n"))
}