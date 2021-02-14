package main

import (
	"npc-generator-be/cmd/app"
	"npc-generator-be/internal/config"
)

func main() {
	// Load config variables before initializing other components
	conf := config.InitializeConfig()

	// Set up database, server
	db := app.InitializeDB(conf)
	app.InitializeRoutes(conf, db)
}
