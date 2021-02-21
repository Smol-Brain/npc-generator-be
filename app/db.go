package app

import (
	"fmt"
	"log"

	generator "github/Smol-Brain/npc-generator-be"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitializeDB sets up database connection
func InitializeDB(config Config) (db *gorm.DB) {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s",
		config.DBUSER,
		config.DBPASSWORD,
		config.DBNAME,
		config.DBHOST,
		config.DBPORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Postgres: " + err.Error())
	}

	log.Println("Connected to Postgres")

	db.AutoMigrate(&generator.Npc{})

	return
}
