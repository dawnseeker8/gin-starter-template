package config

import "github.com/spf13/viper"

// Config is the configuration for the whole gin server
type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DbHost        string `mapstructure:"DB_HOST"`
	DbUser        string `mapstructure:"DB_USER"`
	DbPassword    string `mapstructure:"DB_PASSWORD"`
	DbName        string `mapstructure:"DB_NAME"`
	DbPort        string `mapstructure:"DB_PORT"`
	DbSslMode     string `mapstructure:"DB_SSL_MODE"`
}

// LoadConfig will load the configuration setting from .env file
func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
