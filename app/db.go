package app

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"npc-generator-be/global"
)

// InitializeDB sets up database connection
func InitializeDB() {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		global.Conf.DBUser,
		global.Conf.DBPassword,
		global.Conf.DBName,
		global.Conf.DBHost,
		global.Conf.DBPort,
	)

	var err error
	global.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
