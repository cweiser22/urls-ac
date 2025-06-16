package main

import (
	"github.com/cweiser22/urls-ac/db"
	"github.com/cweiser22/urls-ac/handlers"
	"github.com/cweiser22/urls-ac/metrics"
	"github.com/cweiser22/urls-ac/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	viper.AutomaticEnv()

	viper.BindEnv("postgres_connection_string", "POSTGRES_CONNECTION_STRING")
	viper.BindEnv("redis_connection_string", "REDIS_CONNECTION_STRING")
	viper.BindEnv("port", "PORT")
	viper.BindEnv("environment", "ENVIRONMENT")

	viper.SetDefault("postgres_connection_string", "postgres://postgres:postgres@postgres:5432/url_shortener?sslmode=disable")
	viper.SetDefault("redis_connection_string", "redis:6379")
	viper.SetDefault("port", "8080")
	viper.SetDefault("environment", "dev")

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
