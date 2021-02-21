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
// TODO: Consider replacing with https://github.com/kelseyhightower/envconfig
func InitializeConfig() (config Config) {
	// Heroku has env vars defined instead of a configuration file
	if env := os.Getenv("ENVIRONMENT"); env == "Integration" {
		config.PORT = os.Getenv("PORT")
		config.DBHOST = os.Getenv("DBHOST")
		config.DBNAME = os.Getenv("DBNAME")
		config.DBPASSWORD = os.Getenv("DBPASSWORD")
		config.DBPORT = os.Getenv("DBPORT")
		config.DBUSER = os.Getenv("DBUSER")
	} else {
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
	}

	return
}
