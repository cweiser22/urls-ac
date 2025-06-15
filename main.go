package main

import (
	"github.com/cweiser22/urls-ac/db"
	"github.com/cweiser22/urls-ac/handlers"
	"github.com/cweiser22/urls-ac/metrics"
	"github.com/cweiser22/urls-ac/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	DB, err := db.NewPostgresDB("postgres://postgres:postgres@localhost:5532/url_shortener?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Initialize the Redis client
	// Assuming Redis is running on localhost:6479
	redisClient, err := db.NewRedisClient("localhost:6479")
	if err != nil {
		panic(err)
	}

	shortCodeService := service.NewShortCodeService(DB, redisClient, metrics.CacheRequests)

	indexHandler := handlers.NewIndexHandler()
	healthCheckHandler := handlers.NewHealthCheckHandler()
	urlHandler := handlers.NewURLHandler(shortCodeService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", indexHandler.IndexHandler)

	r.Get("/{shortCode}", urlHandler.RedirectFromMapping)
	r.Get("/api/v1/health", healthCheckHandler.HealthCheckHandler)
	r.Post("/api/v1/mappings", urlHandler.CreateShortURL)

	r.Handle("/api/v1/metrics", promhttp.Handler())

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
