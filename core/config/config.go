package config

import (
	_ "embed"
	"os"
	"strings"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	_ "github.com/joho/godotenv/autoload"
)

//go:embed token_jwt_key.pem
var jwtPublicKey string

// Config is the configuration for the whole gin server
type Config struct {
	ServerAddress       string `default:"0.0.0.0:8080"`
	CasdoorEndpoint     string
	CasdoorClientId     string
	CasdoorClientSecret string
	CasdoorOrganization string
	CasdoorAppName      string

	RedisURL      string
	SessionSecret string
	SessionName   string
}

// LoadConfig will load the configuration setting from .env file
func LoadConfig() (config Config) {

	if os.Getenv("SERVER_ADDRESS") == "" {
		config.ServerAddress = "0.0.0.0:8080"
	} else {
		config.ServerAddress = os.Getenv("SERVER_ADDRESS")
	}

	config.CasdoorEndpoint = strings.TrimRight(os.Getenv("CASDOOR_ENDPOINT"), "/")
	config.CasdoorClientId = os.Getenv("CASDOOR_CLIENT_ID")
	config.CasdoorClientSecret = os.Getenv("CASDOOR_CLIENT_SECRET")
	config.CasdoorOrganization = os.Getenv("CASDOOR_ORGANIZATION")
	config.CasdoorAppName = os.Getenv("CASDOOR_APP_NAME")

	config.RedisURL = os.Getenv("REDIS_URL")
	config.SessionSecret = os.Getenv("SESSION_SECRET")
	config.SessionName = os.Getenv("SESSION_NAME")

	return config
}

func InitAuthConfig(cfg *Config) {

	casdoorsdk.InitConfig(cfg.CasdoorEndpoint, cfg.CasdoorClientId, cfg.CasdoorClientSecret, jwtPublicKey, cfg.CasdoorOrganization, cfg.CasdoorAppName)
}
