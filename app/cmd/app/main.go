package main

import (
	"github.com/cweiser22/urls-ac/internal/config"
	"github.com/cweiser22/urls-ac/internal/db"
	"github.com/cweiser22/urls-ac/internal/handlers"
	"github.com/cweiser22/urls-ac/internal/metrics"
	"github.com/cweiser22/urls-ac/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"

	"net/http"
)

func main() {
	config.Init()
	postgresConnectionString := viper.GetString("postgres_connection_string")
	redisConnectionString := viper.GetString("redis_connection_string")

	DB, err := db.NewPostgresDB(postgresConnectionString)
	if err != nil {
		panic(err)
	}

	// Initialize the Redis client
	// Assuming Redis is running on localhost:6479
	redisClient, err := db.NewRedisClient(redisConnectionString)
	if err != nil {
		panic(err)
	}

	shortCodeService := service.NewShortCodeService(DB, redisClient, metrics.CacheRequestsTotal)

	indexHandler := handlers.NewIndexHandler()
	healthCheckHandler := handlers.NewHealthCheckHandler()
	urlHandler := handlers.NewURLHandler(shortCodeService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/api/v1/metrics", promhttp.Handler())
	r.Get("/", indexHandler.AppHandler)
	r.Get("/{shortCode}", urlHandler.RedirectFromMapping)
	r.Get("/api/v1/health", healthCheckHandler.HealthCheckHandler)
	r.Post("/api/v1/mappings", urlHandler.CreateShortURL)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
