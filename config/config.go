package config

import (
	"github.com/spf13/viper"
)

// Config holds environment and file defined variables
type Config struct {
	// Server config
	Port string

	// Database config
	DBHost     string
	DBName     string
	DBPassword string
	DBPort     string
	DBUser     string
}

// InitializeConfig loads configuration values
func InitializeConfig() (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	// Allow env vars to take precedence
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return
}
