package main

import (
	"github.com/cweiser22/urls-ac/internal/cache"
	"github.com/cweiser22/urls-ac/internal/config"
	"github.com/cweiser22/urls-ac/internal/db"
	"github.com/cweiser22/urls-ac/internal/handlers"
	"github.com/cweiser22/urls-ac/internal/metrics"
	"github.com/cweiser22/urls-ac/internal/repository"
	"github.com/cweiser22/urls-ac/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"

	"net/http"
)

func main() {
	config.Init()
	postgresConnectionString := viper.GetString("postgres_connection_string")
	redisConnectionString := viper.GetString("redis_connection_string")
	environment := viper.GetString("environment")

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

	urlMappingCache := cache.NewURLMappingCache(redisClient)

	urlMappingRepository := repository.NewURLMappingsRepository(DB)
	fiftyFiftyRepository := repository.NewFiftyFiftyLinkRepository(DB)

	shortenService := service.NewShortenService(urlMappingCache, metrics.CacheRequestsTotal, urlMappingRepository)
	shortCodeService := service.NewShortCodeService()
	fiftyFiftyService := service.NewFiftyFiftyLinkService(fiftyFiftyRepository)

	indexHandler := handlers.NewIndexHandler()
	healthCheckHandler := handlers.NewHealthCheckHandler()
	urlHandler := handlers.NewURLHandler(shortenService)
	fiftyFiftyHandler := handlers.NewFiftyFiftyHandler(fiftyFiftyService, shortCodeService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}

	if environment == "production" {
		corsOptions.AllowedOrigins = []string{"https://urls.ac"}
	}

	r.Use(cors.Handler(corsOptions))

	// prometheus metrics endpoint
	r.Handle("/api/v1/metrics", promhttp.Handler())

	r.Get("/", indexHandler.AppHandler)

	// Redirect from short code to original URL
	r.Get("/{shortCode}", urlHandler.RedirectFromMapping)

	r.Get("/ff/{shortCode}", fiftyFiftyHandler.Redirect)
	r.Post("/api/v1/ff/", fiftyFiftyHandler.Create)

	r.Get("/api/v1/health", healthCheckHandler.HealthCheckHandler)
	r.Post("/api/v1/mappings", urlHandler.CreateShortURL)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
