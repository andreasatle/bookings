package main

import (
	"net/http"

	"github.com/andreasatle/bookings/config"
	"github.com/andreasatle/bookings/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// Create a new router
	mux := chi.NewRouter()

	// Setup middleware
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(LoadSession)

	// Route traffic
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about/", handlers.Repo.About)

	return mux
}
