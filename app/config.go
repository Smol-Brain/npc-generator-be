package app

import (
	"github.com/spf13/viper"
)

// Config holds environment and file defined variables
type Config struct {
	// Server config
	PORT string

	// Database config
	DBHOST     string
	DBNAME     string
	DBPASSWORD string
	DBPORT     string
	DBUSER     string
}

// InitializeConfig loads configuration values
func InitializeConfig() (config Config) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	// Allow env vars to take precedence
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read configuration file: " + err.Error())
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic("Failed to unmarshal configuration: " + err.Error())
	}

	return
}
