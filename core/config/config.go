package config

import "os"

// Config is the configuration for the whole gin server
type Config struct {
	ServerAddress string `default:"0.0.0.0:8080"`
	DbHost        string `mapstructure:"DB_HOST"`
	DbUser        string `mapstructure:"DB_USER"`
	DbPassword    string `mapstructure:"DB_PASSWORD"`
	DbName        string `mapstructure:"DB_NAME"`
	DbPort        string `mapstructure:"DB_PORT"`
	DbSslMode     string `mapstructure:"DB_SSL_MODE"`
}

// LoadConfig will load the configuration setting from .env file
func LoadConfig() (config Config) {

	if os.Getenv("SERVER_ADDRESS") == "" {
		config.ServerAddress = "0.0.0.0:8080"
	} else {
		config.ServerAddress = os.Getenv("SERVER_ADDRESS")
	}
	return config
}
