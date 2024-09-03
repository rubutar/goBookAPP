package main

import (
	"net/http"

	"github.com/rubutar/goBookAPP/internal/config"
	"github.com/rubutar/goBookAPP/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals", handlers.Repo.Generals)
	mux.Get("/majors", handlers.Repo.Majors)
	mux.Get("/reserve", handlers.Repo.Reserve)
	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/availability", handlers.Repo.Availability)
	mux.Post("/availability", handlers.Repo.PostAvailability)
	mux.Post("/availability-json", handlers.Repo.AvailabilityJSON)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
