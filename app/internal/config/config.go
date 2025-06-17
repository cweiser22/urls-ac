package config

import "github.com/spf13/viper"

func Init() {
	viper.AutomaticEnv()

	viper.BindEnv("postgres_connection_string", "POSTGRES_CONNECTION_STRING")
	viper.BindEnv("redis_connection_string", "REDIS_CONNECTION_STRING")
	viper.BindEnv("port", "PORT")
	viper.BindEnv("environment", "ENVIRONMENT")
	viper.BindEnv("host", "HOST")

	viper.SetDefault("postgres_connection_string", "postgres://postgres:postgres@postgres:5432/url_shortener?sslmode=disable")
	viper.SetDefault("redis_connection_string", "redis:6379")
	viper.SetDefault("port", "8080")
	viper.SetDefault("environment", "production")
	viper.SetDefault("host", "localhost")
}
