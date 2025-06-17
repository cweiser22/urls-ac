package main

import (
	"github.com/cweiser22/urls-ac/app/internal/config"
	db2 "github.com/cweiser22/urls-ac/app/internal/db"
	handlers2 "github.com/cweiser22/urls-ac/app/internal/handlers"
	"github.com/cweiser22/urls-ac/app/internal/metrics"
	"github.com/cweiser22/urls-ac/app/internal/service"
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

	DB, err := db2.NewPostgresDB(postgresConnectionString)
	if err != nil {
		panic(err)
	}

	// Initialize the Redis client
	// Assuming Redis is running on localhost:6479
	redisClient, err := db2.NewRedisClient(redisConnectionString)
	if err != nil {
		panic(err)
	}

	shortCodeService := service.NewShortCodeService(DB, redisClient, metrics.CacheRequestsTotal)

	indexHandler := handlers2.NewIndexHandler()
	healthCheckHandler := handlers2.NewHealthCheckHandler()
	urlHandler := handlers2.NewURLHandler(shortCodeService)

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
