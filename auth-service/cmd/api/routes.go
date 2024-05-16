package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (c *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// health check
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/authenticate", c.Authenticate)

	return mux
}
