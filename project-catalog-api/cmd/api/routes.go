package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/v1/healthcheck", http.HandlerFunc(app.healthCheckHandler))
	router.Handler(http.MethodGet, "/v1/projects/:id", http.HandlerFunc(app.showProjectHandler))

	return router
}
