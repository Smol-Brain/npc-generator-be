package app

import (
	"os"

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

	// Hacky, overriding with env vars for now to deal with Heroku issue
	if port := os.Getenv("PORT"); port != "" {
		config.PORT = port
	}

	if dbHost := os.Getenv("DBHOST"); dbHost != "" {
		config.DBHOST = dbHost
	}

	if dbName := os.Getenv("DBNAME"); dbName != "" {
		config.DBNAME = dbName
	}

	if dbPassword := os.Getenv("DBPASSWORD"); dbPassword != "" {
		config.DBPASSWORD = dbPassword
	}

	if dbPort := os.Getenv("DBPORT"); dbPort != "" {
		config.DBPORT = dbPort
	}

	if dbUser := os.Getenv("DBUSER"); dbUser != "" {
		config.DBUSER = dbUser
	}

	return
}
